/*

Copyright (c) 2018 - 2024 PhotoPrism UG. All rights reserved.

    This program is free software: you can redistribute it and/or modify
    it under Version 3 of the GNU Affero General Public License (the "AGPL"):
    <https://docs.photoprism.app/license/agpl>

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    The AGPL is supplemented by our Trademark and Brand Guidelines,
    which describe how our Brand Assets may be used:
    <https://www.photoprism.app/trademark>

Feel free to send an email to hello@photoprism.app if you have questions,
want to support our work, or just want to say hello.

Additional information can be found in our Developer Guide:
<https://docs.photoprism.app/developer-guide/>

*/

import PhotoSwipe from "photoswipe";
import LightBox from "photoswipe/lightbox";
import Event from "pubsub-js";
import Util from "util.js";
import Api from "./api";
import Thumb from "model/thumb";

const thumbs = window.__CONFIG__.thumbs;

class Viewer {
  constructor() {
    this.el = null;
    this.lightbox = null;
  }

  getEl() {
    if (!this.el) {
      this.el = document.getElementById("photo-viewer");

      if (this.el === null) {
        let err = "no photo viewer element found";
        console.warn(err);
        throw err;
      }
    }

    return this.el;
  }

  play(params) {
    Event.publish("player.open", params);
  }

  show(items, index = 0) {
    if (!Array.isArray(items) || items.length === 0 || index >= items.length) {
      console.log("photo list passed to gallery was empty:", items);
      return;
    }

    // PhotoSwipe configuration options, see https://photoswipe.com/options/.
    const options = {
      pswpModule: PhotoSwipe,
      dataSource: items,
      index: index,
      mouseMovePan: true,
      arrowPrev: true,
      arrowNext: true,
      zoom: true,
      close: true,
      counter: true,
      initialZoomLevel: "fit",
      secondaryZoomLevel: "fill",
      maxZoomLevel: 3,
      bgOpacity: 1,
      preload: [1, 1],
      returnFocus: true,
      showHideAnimationType: "none",
      tapAction: "toggle-controls",
      imageClickAction: "toggle-controls",
    };

    // Create PhotoSwipe instance.
    let lightbox = new LightBox(options);
    let firstPicture = true;

    // Keep reference to PhotoSwipe instance.
    this.lightbox = lightbox;

    Event.publish("viewer.show");

    // Add a close event handler to destroy the viewer after use.
    lightbox.on("close", () => {
      this.lightbox.destroy();
      this.lightbox = null;
      Event.publish("viewer.pause");
      Event.publish("viewer.hide");
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

      // Picture caption displayed at the bottom.
      lightbox.pswp.ui.registerElement({
        name: "caption",
        order: 9,
        isButton: false,
        appendTo: "root",
        html: "",
        onInit: (el, pswp) => {
          lightbox.pswp.on("change", () => {
            const data = items[lightbox.pswp.currSlide.index];

            let caption = "";

            if (data.Title) {
              caption += `<h4>${Util.encodeHTML(data.Title)}</h4>`;
            }

            if (data.Description) {
              caption += `<p>${Util.encodeHTML(data.Description)}</p>`;
            }

            if (data.Playable) {
              el.classList.add("pswp__caption-video");
            } else {
              el.classList.remove("pswp__caption-video");
            }

            el.innerHTML = Util.sanitizeHtml(caption);
          });
        },
      });
    });

    // Process raw data for PhotoSwipe, see https://photoswipe.com/filters/#itemdata.
    //
    // Todo: Should be improved to allow dynamic zooming and play videos in their native format whenever possible.
    lightbox.addFilter("itemData", (el, i) => {
      const data = items[i];
      const viewportWidth = window.innerWidth * window.devicePixelRatio;
      const viewportHeight = window.innerHeight * window.devicePixelRatio;

      const s = Util.thumbSize(viewportWidth, viewportHeight);

      const imgSrc = data.Thumbs[s].src;

      if (data.Playable) {
        const videoSrc = Util.videoUrl(data.Hash);
        if (firstPicture) {
          firstPicture = false;
          return {
            html: `<video class="pswp__video" autoplay controls playsinline poster="${imgSrc}" style="width: 100vw; height: 100vh" preload="auto"><source src="${videoSrc}" /></video>`,
          };
        } else {
          return {
            html: `<video class="pswp__video" controls playsinline poster="${imgSrc}" style="width: 100vw; height: 100vh" preload="auto"><source src="${videoSrc}" /></video>`,
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
  }

  // Loads picture data and then opens the viewer.
  static show(ctx, index) {
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
        ctx.$viewer.show(ctx.viewer.results, i);
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

        // Show photos.
        ctx.$viewer.show(ctx.viewer.results, i);
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
  }
}

export default Viewer;
