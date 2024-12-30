<template>
  <v-dialog :model-value="show" max-width="380">
    <v-card>
      <v-card-title class="d-flex justify-start align-center flex-nowrap ga-3 text-h6">
        <v-icon icon="mdi-alert-decagram-outline" color="primary"></v-icon>
        <h6 class="text-h6"><translate>Software Update</translate></h6>
      </v-card-title>
      <v-card-text class="d-flex justify-start flex-column ga-3">
        <div class="text-body-2 data-message">{{ getMessage() }}</div>
        <div dir="ltr" class="text-caption data-version">
          {{ getVersion() }}
        </div>
      </v-card-text>
      <v-card-actions>
        <v-btn color="button" variant="flat" @click="close">
          <translate>Close</translate>
        </v-btn>
        <v-btn color="highlight" class="action-update-reload" variant="flat" @click="reload">
          <translate>Reload</translate>
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "PReloadDialog",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      visible: this.show,
    };
  },
  watch: {
    show(val) {
      this.visible = val;
    },
    visible(val) {
      if (!val) {
        this.close();
      }
    },
  },
  methods: {
    getMessage() {
      return this.$gettext("A new version of %{s} is available:", { s: this.$config.getAbout() });
    },
    getVersion() {
      return this.$config.getServerVersion();
    },
    close() {
      this.$emit("close");
    },
    reload() {
      this.$notify.info(this.$gettext("Reloading…"));
      this.$notify.blockUI();
      setTimeout(() => window.location.reload(), 100);
    },
  },
};
</script>
