<template>
  <v-dialog :model-value="show" persistent max-width="540" class="p-share-dialog" @keydown.esc="close">
    <v-card>
      <v-card-title class="d-flex justify-space-between align-center ga-3">
        <h6 class="text-h6"><translate :translate-params="{ name: model.modelName() }">Share %{name}</translate></h6>
        <v-btn icon="mdi-link-plus" variant="text" color="primary" :title="$gettext('Add Link')" class="action-add-link" @click.stop="add"></v-btn>
      </v-card-title>
      <v-card-text>
        <v-expansion-panels variant="accordion" density="compact" rounded="6" class="elevation-0">
          <v-expansion-panel v-for="(link, index) in links" :key="link.UID" color="secondary" class="pa-0 elevation-0">
            <v-expansion-panel-title>
              <button class="text-start action-url ml-0 mt-0 mb-0 pa-0" style="user-select: none" @click.stop="copyUrl(link)">
                <v-icon size="16" class="pr-1">mdi-link</v-icon>
                /s/<strong v-if="link.Token" style="font-weight: 500"> {{ link.getToken() }} </strong><span v-else>…</span>
              </button>
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <v-card color="secondary-light">
                <v-card-text class="dense">
                  <v-row dense>
                    <v-col cols="12">
                      <v-text-field :model-value="link.url()" hide-details density="comfortable" variant="solo" flat readonly :label="$gettext('URL')" autocorrect="off" autocapitalize="none" autocomplete="off" class="input-url" @click.stop="selectText($event)"> </v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-select v-model="link.Expires" hide-details density="comfortable" variant="solo" flat :label="expires(link)" browser-autocomplete="off" :items="options.Expires()" item-title="text" item-value="value" class="input-expires"> </v-select>
                    </v-col>
                    <v-col cols="12" sm="6">
                      <v-text-field v-model="link.Token" hide-details density="comfortable" variant="solo" flat required autocomplete="off" autocorrect="off" autocapitalize="none" :label="$gettext('Secret')" :placeholder="$gettext('Token')" class="input-secret"></v-text-field>
                    </v-col>
                    <!-- <v-col cols="12" sm="6" class="pa-2">
                      <v-text-field
                        v-model="link.Password"
                        hide-details
                        autocomplete="off"
                        :label="label.pass"
                        :placeholder="link.HasPassword ? '••••••••' : 'optional'"
                        color="surface-variant"
                        :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                        :type="showPassword ? 'text' : 'password'"
                        @click:append-inner="showPassword = !showPassword"
                      ></v-text-field>
                    </v-col> -->
                    <v-col cols="12" class="d-flex justify-space-between align-center ga-3">
                      <v-btn variant="text" color="remove" density="comfortable" icon="mdi-delete" :title="$gettext('Delete')" class="action-delete" @click.stop.exact="remove(index)"> </v-btn>
                      <v-btn variant="flat" color="highlight" class="action-save" @click.stop.exact="update(link)">
                        <translate>Save</translate>
                      </v-btn>
                    </v-col>
                  </v-row>
                </v-card-text>
              </v-card>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>

        <div class="pt-3 text-caption">
          <translate :translate-params="{ name: model.modelName() }">People you share a link with will be able to view public contents.</translate>
          <translate>A click will copy it to your clipboard.</translate>
          <translate>Any private photos and videos remain private and won't be shared.</translate>
          <translate>Alternatively, you can upload files directly to WebDAV servers like Nextcloud.</translate>
        </div>
      </v-card-text>
      <v-card-actions>
        <v-btn variant="flat" color="button" class="action-webdav" @click.stop="upload">
          <translate>WebDAV Upload</translate>
        </v-btn>
        <v-btn variant="flat" color="button" class="action-close" @click.stop="confirm">
          <translate>Close</translate>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
import * as options from "options/options";
import Util from "common/util";

export default {
  name: "PShareDialog",
  props: {
    show: Boolean,
    model: {
      type: Object,
      default: () => {},
    },
  },
  data() {
    return {
      host: window.location.host,
      showPassword: false,
      loading: false,
      search: null,
      links: [],
      options: options,
      label: {
        url: this.$gettext("Service URL"),
        user: this.$gettext("Username"),
        pass: this.$gettext("Password"),
        cancel: this.$gettext("Cancel"),
        confirm: this.$gettext("Done"),
      },
      rtl: this.$rtl,
    };
  },
  watch: {
    show: function (show) {
      if (show) {
        this.links = [];
        this.loading = true;
        this.model
          .links()
          .then((resp) => {
            if (resp.count === 0) {
              this.add();
            } else {
              this.links = resp.models;
            }
          })
          .finally(() => (this.loading = false));
      }
    },
  },
  methods: {
    selectText(ev) {
      if (!ev || !ev.target) {
        return;
      }

      ev.target.select();
    },
    async copyUrl(link) {
      try {
        const url = link.url();
        await Util.copyToMachineClipboard(url);
        this.$notify.success(this.$gettext("Copied to clipboard"));
      } catch (error) {
        this.$notify.error(this.$gettext("Failed copying to clipboard"));
      }
    },
    expires(link) {
      let result = this.$gettext("Expires");

      if (link.Expires <= 0) {
        return result;
      }

      return `${result}: ${link.expires()}`;
    },
    add() {
      this.loading = true;

      this.model
        .createLink()
        .then((r) => {
          this.links.push(r);
        })
        .finally(() => (this.loading = false));
    },
    update(link) {
      if (!link) {
        this.$notify.error(this.$gettext("Failed updating link"));
        return;
      }

      this.loading = true;

      this.model
        .updateLink(link)
        .then(() => {
          this.$notify.success(this.$gettext("Changes successfully saved"));
        })
        .finally(() => {
          this.loading = false;
        });
    },
    remove(index) {
      const link = this.links[index];

      if (!link) {
        this.$notify.error(this.$gettext("Failed removing link"));
        return;
      }

      this.loading = true;

      this.model
        .removeLink(link)
        .then(() => {
          this.$notify.success(this.$gettext("Changes successfully saved"));
          this.links.splice(index, 1);
        })
        .finally(() => (this.loading = false));
    },
    upload() {
      this.$emit("upload");
    },
    close() {
      this.$emit("close");
    },
    confirm() {
      this.$emit("close");
    },
  },
};
</script>
