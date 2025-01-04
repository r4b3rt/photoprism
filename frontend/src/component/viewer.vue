<template>
  <div v-if="visible" ref="container" class="p-viewer" tabindex="-1" role="dialog">
    <!-- div class="pswp__bg"></div>
    <div class="pswp__scroll-wrap">
      <div class="pswp__container" :class="{ slideshow: slideshow.active }">
        <div class="pswp__item"></div>
        <div class="pswp__item"></div>
        <div class="pswp__item"></div>
      </div>

      <div class="pswp__ui pswp__ui--hidden">
        <div class="pswp__top-bar">
          <div class="pswp__taken hidden-xs">
            {{ formatDate(slide.TakenAtLocal) }}
          </div>

          <div class="pswp__counter"></div>

          <button class="pswp__button pswp__button--close action-close" :title="$gettext('Close')"></button>

          <button v-if="canDownload" class="pswp__button action-download" style="background: none" :title="$gettext('Download')" @click.exact="onDownload">
            <v-icon size="16" color="white">mdi-download</v-icon>
          </button>

          <button v-if="canEdit" class="pswp__button action-edit hidden-shared-only" style="background: none" :title="$gettext('Edit')" @click.exact="onEdit">
            <v-icon size="16" color="white">mdi-pencil</v-icon>
          </button>

          <button class="pswp__button action-select" style="background: none" :title="$gettext('Select')" @click.exact="onSelect">
            <v-icon v-if="selection.length && $clipboard.has(item)" size="16" color="white">mdi-check-circle</v-icon>
            <v-icon v-else size="16" color="white">mdi-circle-outline</v-icon>
          </button>

          <button v-if="canLike" class="pswp__button action-like hidden-shared-only" style="background: none" :title="$gettext('Like')" @click.exact="onLike">
            <v-icon v-if="slide.Favorite" icon="mdi-star" size="19" color="white"></v-icon>
            <v-icon v-else icon="mdi-star-outline" size="19" color="white"></v-icon>
          </button>

          <button class="pswp__button pswp__button--fs action-toggle-fullscreen" :title="$gettext('Fullscreen')"></button>

          <button class="pswp__button pswp__button--zoom action-zoom" :title="$gettext('Zoom in/out')"></button>

          <button class="pswp__button action-slideshow" style="background: none" :title="$gettext('Start/Stop Slideshow')" @click.exact="onSlideshow">
            <v-icon v-show="!interval" size="18" color="white">mdi-play</v-icon>
            <v-icon v-show="interval" size="16" color="white">mdi-pause</v-icon>
          </button>

          <div class="pswp__preloader">
            <div class="pswp__preloader__icn">
              <div class="pswp__preloader__cut">
                <div class="pswp__preloader__donut"></div>
              </div>
            </div>
          </div>
        </div>

        <div class="pswp__share-modal pswp__share-modal--hidden pswp__single-tap">
          <div class="pswp__share-tooltip"></div>
        </div>

        <button class="pswp__button pswp__button--arrow--left action-previous" title="Previous (arrow left)"></button>

        <button class="pswp__button pswp__button--arrow--right action-next" title="Next (arrow right)"></button>

        <div class="pswp__caption" @click="onPlay">
          <div class="pswp__caption__center"></div>
        </div>
      </div>
    </div -->
  </div>
</template>

<script>
import PhotoSwipe from "photoswipe";
import LightBox from "photoswipe/lightbox";
import PhotoSwipeDynamicCaption from "photoswipe-dynamic-caption-plugin";
import Util from "common/util";
import Api from "common/api";
import Thumb from "model/thumb";
import { Photo } from "model/photo";
import Notify from "common/notify";

/*
  TODO: All previously available features and controls must be preserved in the new hybrid photo/video viewer:
    1. Some of the controls that the old viewer had (e.g. (a) select, (b) play slideshow, (c) fullscreen,
       (d) edit, (e) date info,...) are still missing.
    2. The existing controls need improvements (e.g. (a) the download button currently only downloads the
       thumbnail instead of having the same functionality as before, (b) the zoom doesn't load a larger
       version of the image yet).
    3. Finally, after the refactoring/upgrade, (a) the old/unused code (e.g. for the separate video player) needs
       to be removed and (b) everything needs to be thoroughly tested on all major browsers and mobile devices.
*/
export default {
  name: "PViewer",
  data() {
    return {
      visible: false,
      lightbox: null,
      captionPlugin: null,
      captionTimer: false,
      hasTouch: false,
      idleTime: 4000,
      canEdit: this.$config.allow("photos", "update") && this.$config.feature("edit"),
      canLike: this.$config.allow("photos", "manage") && this.$config.feature("favorites"),
      canDownload: this.$config.allow("photos", "download") && this.$config.feature("download"),
      selection: this.$clipboard.selection,
      config: this.$config.values,
      slide: new Thumb(),
      subscriptions: [],
      interval: false,
      slideshow: {
        active: false,
        next: 0,
      },
      player: {
        show: false,
        loop: false,
        autoplay: true,
        source: "",
        poster: "",
        width: 640,
        height: 480,
      },
    };
  },
  created() {
    // this.subscriptions["viewer.change"] = this.$event.subscribe("viewer.change", this.onChange);
    this.subscriptions["viewer.pause"] = this.$event.subscribe("viewer.pause", this.onPause);
    // this.subscriptions["viewer.show"] = this.$event.subscribe("viewer.show", this.onShow);
    this.subscriptions["viewer.hide"] = this.$event.subscribe("viewer.hide", this.onHide);
  },
  unmounted() {
    this.onPause();

    this.lightbox.destroy();
    this.lightbox = null;

    for (let i = 0; i < this.subscriptions.length; i++) {
      this.$event.unsubscribe(this.subscriptions[i]);
    }
  },
  methods: {
    getEl() {
      return this.$refs?.container;
    },
    getPhotoSwipe() {
      return this.lightbox?.pswp;
    },
    getVideos() {
      return this.getEl().getElementsByTagName("video");
    },
    showContext(ctx, index) {
      if (ctx.loading || !ctx.listen || ctx.viewer.loading || !ctx.results[index]) {
        return false;
      }

      const selected = ctx.results[index];

      if (!ctx.viewer.dirty && ctx.viewer.results && ctx.viewer.results.length > index) {
        // Reuse existing viewer result if possible.
        let i = -1;

        if (ctx.viewer.results[index] && ctx.viewer.results[index].UID === selected.UID) {
          i = index;
        } else {
          i = ctx.viewer.results.findIndex((p) => p.UID === selected.UID);
        }

        if (i > -1 && (((ctx.viewer.complete || ctx.complete) && ctx.viewer.results.length >= ctx.results.length) || i + ctx.viewer.batchSize <= ctx.viewer.results.length)) {
          this.showThumbs(ctx.viewer.results, i);
          return;
        }
      }

      // Fetch photos from server API.
      ctx.viewer.loading = true;

      const params = ctx.searchParams();
      params.count = params.offset + ctx.viewer.batchSize;
      params.offset = 0;

      // Fetch viewer results from API.
      return Api.get("photos/view", { params })
        .then((response) => {
          const count = response && response.data ? response.data.length : 0;
          if (count === 0) {
            ctx.$notify.warn(ctx.$gettext("No pictures found"));
            ctx.viewer.dirty = true;
            ctx.viewer.complete = false;
            return;
          }

          // Process response.
          if (response.headers && response.headers["x-count"]) {
            const c = parseInt(response.headers["x-count"]);
            const l = parseInt(response.headers["x-limit"]);
            ctx.viewer.complete = c < l;
          } else {
            ctx.viewer.complete = ctx.complete;
          }

          let i;

          if (response.data[index] && response.data[index].UID === selected.UID) {
            i = index;
          } else {
            i = response.data.findIndex((p) => p.UID === selected.UID);
          }

          ctx.viewer.results = Thumb.wrap(response.data);

          // Show pictures.
          this.showThumbs(ctx.viewer.results, i);
          ctx.viewer.dirty = false;
        })
        .catch(() => {
          ctx.viewer.dirty = true;
          ctx.viewer.complete = false;
        })
        .finally(() => {
          // Unblock.
          ctx.viewer.loading = false;
        });
    },
    showThumbs(slides, index = 0) {
      if (!Array.isArray(slides) || slides.length === 0 || index >= slides.length) {
        console.log("item list passed to viewer is empty:", slides);
        return;
      }

      this.onShow();

      this.$nextTick(() => {
        this.renderLightBox(slides, index);
      });
    },
    renderLightBox(slides, index = 0) {
      if (!Array.isArray(slides) || slides.length === 0 || index >= slides.length) {
        console.log("item list passed to viewer is empty:", slides);
        return;
      }

      this.$event.publish("viewer.show");

      // PhotoSwipe configuration options, see https://photoswipe.com/options/.
      const options = {
        appendToEl: this.getEl(),
        pswpModule: PhotoSwipe,
        dataSource: slides,
        index: index,
        mouseMovePan: true,
        arrowPrev: true,
        arrowNext: true,
        zoom: true,
        close: true,
        counter: false,
        initialZoomLevel: "fit",
        secondaryZoomLevel: "fill",
        maxZoomLevel: 3,
        bgOpacity: 1,
        preload: [1, 1],
        showHideAnimationType: "none",
        tapAction: "toggle-controls",
        imageClickAction: "zoom",
        mainClass: "media-viewer-lightbox",
        paddingFn: () => {
          return {
            top: 24,
            bottom: 24,
            left: 0,
            right: 0,
          };
        },
      };

      // Create PhotoSwipe instance.
      let lightbox = new LightBox(options);
      let firstPicture = true;

      // Keep reference to PhotoSwipe instance.
      this.lightbox = lightbox;
      this.captionTimer = false;
      this.hasTouch = false;

      // Add dynamic caption plugin.
      this.captionPlugin = new PhotoSwipeDynamicCaption(lightbox, {
        type: "auto",
        captionContent: (slide) => {
          if (!slide || !slides || slide?.index < 0) {
            return "";
          }

          const media = slides[slide.index];

          if (media) {
            return this.formatCaption(media);
          }

          return "";
        },
      });

      // Add a close event handler to destroy the viewer after use.
      lightbox.on("close", () => {
        this.$event.publish("viewer.pause");
        this.$event.publish("viewer.hide");
      });

      // Add user interface controls.
      //
      // Todo: The same controls as with PhotoSwipe 4 should be usable/available!
      lightbox.on("uiRegister", function () {
        // Download button displayed at the top.
        // Todo: Proof-of-concept, requires refactoring.
        lightbox.pswp.ui.registerElement({
          name: "download-button",
          order: 8,
          isButton: true,
          tagName: "a",

          // SVG with outline
          html: {
            isCustomSVG: true,
            inner: '<path d="M20.5 14.3 17.1 18V10h-2.2v7.9l-3.4-3.6L10 16l6 6.1 6-6.1ZM23 23H9v2h14Z" id="pswp__icn-download"/>',
            outlineID: "pswp__icn-download",
          },

          onInit: (el, pswp) => {
            el.setAttribute("download", "");
            el.setAttribute("target", "_blank");
            el.setAttribute("rel", "noopener");

            pswp.on("change", () => {
              el.href = pswp.currSlide.data.src;
            });
          },
        });
      });

      // Auto hide PhotoSwipe UI.
      this.lightbox.on("change", () => {
        this.onChange(slides);
      });

      this.lightbox.on("destroy", () => {
        this.stopHideTimer();
      });

      // Process raw data for PhotoSwipe, see https://photoswipe.com/filters/#itemdata.
      //
      // Todo: Should be improved to allow dynamic zooming and play videos in their native format whenever possible.
      lightbox.addFilter("itemData", (el, i) => {
        const data = slides[i];
        const viewportWidth = window.innerWidth * window.devicePixelRatio;
        const viewportHeight = window.innerHeight * window.devicePixelRatio;

        const s = Util.thumbSize(viewportWidth, viewportHeight);

        const imgSrc = data.Thumbs[s].src;

        if (data.Playable) {
          const videoSrc = Util.videoUrl(data.Hash);
          if (firstPicture) {
            firstPicture = false;
            return {
              html: `<video class="pswp__video" autoplay controls playsinline poster="${imgSrc}" preload="auto"><source src="${videoSrc}" /></video>`,
            };
          } else {
            return {
              html: `<video class="pswp__video" controls playsinline poster="${imgSrc}" preload="auto"><source src="${videoSrc}" /></video>`,
            };
          }
        }

        el.src = imgSrc;
        el.w = Number(data.Thumbs[s].w);
        el.h = Number(data.Thumbs[s].h);

        if (firstPicture) {
          firstPicture = false;
        }

        return el;
      });

      // Init PhotoSwipe.
      lightbox.init();

      // Show first image.
      lightbox.loadAndOpen(index);
    },
    formatCaption(media) {
      if (!media) {
        return "";
      }

      let caption = "";

      if (media.Title) {
        caption += `<h4>${Util.encodeHTML(media.Title)}</h4>`;
      }

      // TODO: Find a good position and styles for the date information.
      /* if (media.TakenAtLocal) {
        caption += `<div>${Util.formatDate(media.TakenAtLocal)}</div>`;
      } */

      if (media.Caption) {
        caption += `<p>${Util.encodeHTML(media.Caption)}</p>`;
      } else if (media.Description) {
        caption += `<p>${Util.encodeHTML(media.Description)}</p>`;
      }

      return Util.sanitizeHtml(caption);
    },
    onShow() {
      this.$scrollbar.hide();
      this.visible = true;
    },
    onHide() {
      this.pauseVideos();
      this.lightbox.destroy();
      this.lightbox = null;
      this.slide = new Thumb();
      this.visible = false;
      this.$scrollbar.show();
    },
    onChange(slides) {
      const pswp = this.getPhotoSwipe();

      if (!pswp) {
        return;
      }

      this.pauseVideos();

      document.addEventListener(
        "touchstart",
        () => {
          this.stopHideTimer();
          this.hasTouch = true;
        },
        { once: true }
      );
      document.addEventListener(
        "mousemove",
        () => {
          this.startHideTimer();
        },
        { once: true }
      );

      if (this.slideshow.next !== pswp.currIndex) {
        this.onPause();
      }

      if (pswp.currIndex && slides && pswp.currIndex >= 0 && pswp.currIndex < slides.length) {
        this.slide = slides[pswp.currIndex];
      }
    },
    onLike() {
      this.slide.toggleLike();
    },
    onSelect() {
      this.$clipboard.toggle(this.slide);
    },
    onPlay() {
      if (this.slide && this.slide.Playable) {
        new Photo().find(this.slide.UID).then((video) => this.openPlayer(video));
      }
    },
    openPlayer(video) {
      if (!video) {
        this.$notify.error(this.$gettext("No video selected"));
        return;
      }

      const params = video.videoParams();

      if (params.error) {
        this.$notify.error(params.error);
        return;
      }

      // Set video parameters.
      this.player.loop = params.loop;
      this.player.width = params.width;
      this.player.height = params.height;
      this.player.poster = params.poster;
      this.player.source = params.uri;

      // Play video.
      this.player.show = true;
    },
    onPause() {
      this.slideshow.active = false;

      if (this.interval) {
        clearInterval(this.interval);
        this.interval = false;
      }
    },
    pauseVideos() {
      const videos = this.getVideos();

      if (!videos || !videos.length) {
        return false;
      }

      for (let video of videos) {
        if (typeof video.pause === "function") {
          try {
            if (!video.paused) {
              video.pause();
            }
          } catch (e) {
            console.log(e);
          }
        }
      }
    },
    onSlideshow() {
      if (this.interval) {
        this.onPause();
        return;
      }

      this.slideshow.active = true;

      const self = this;
      const psp = this.getPhotoSwipe();

      self.interval = setInterval(() => {
        if (psp && typeof psp.next === "function") {
          psp.next();
          this.slideshow.next = psp.currIndex;
        } else {
          this.onPause();
        }
      }, 5000);
    },
    onDownload() {
      this.onPause();

      if (!this.slide || !this.slide.DownloadUrl) {
        console.warn("photo viewer: no download url");
        return;
      }

      Notify.success(this.$gettext("Downloading…"));

      new Photo().find(this.slide.UID).then((p) => p.downloadAll());
    },
    onEdit() {
      this.onPause();

      const g = this.getPhotoSwipe();
      let index = 0;

      // remove duplicates
      let filtered = g.items.filter(function (p, i, s) {
        return !(i > 0 && p.UID === s[i - 1].UID);
      });

      let selection = filtered.map((p, i) => {
        if (g.currItem.UID === p.UID) {
          index = i;
        }

        return p.UID;
      });

      let album = null;

      g.close(); // Close Gallery

      this.$event.publish("dialog.edit", { selection, album, index }); // Open Edit Dialog
    },
    showUI() {
      if (this.getPhotoSwipe() && this.getPhotoSwipe().element) {
        this.getPhotoSwipe().element.classList.add("pswp--ui-visible");
      }
    },
    hideUI() {
      if (this.getPhotoSwipe() && this.getPhotoSwipe().element) {
        this.getPhotoSwipe().element.classList.remove("pswp--ui-visible");
      }
    },
    mouseMove() {
      this.stopHideTimer();
      if (this.lightbox) {
        this.showUI();
        this.startHideTimer();
      }
    },
    startHideTimer() {
      if (this.hasTouch) {
        return;
      }

      this.stopHideTimer();
      this.captionTimer = window.setTimeout(() => {
        this.hideUI();
      }, this.idleTime);
      document.addEventListener(
        "mousemove",
        () => {
          this.mouseMove();
        },
        { once: true }
      );
    },
    stopHideTimer() {
      if (this.captionTimer) {
        window.clearTimeout(this.captionTimer);
        this.captionTimer = false;
      }
    },
  },
};
</script>
