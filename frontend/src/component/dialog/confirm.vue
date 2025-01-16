<template>
  <v-dialog :model-value="show" persistent max-width="350" class="p-confirm-dialog" @keydown.esc="cancel">
    <v-card>
      <v-card-title class="d-flex justify-start align-center ga-3">
        <v-icon size="54" color="primary">{{ icon }}</v-icon>
        <p class="text-subtitle-1">{{ $gettext(`Are you sure?`) }}</p>
      </v-card-title>
      <v-card-actions class="action-buttons">
        <v-btn variant="flat" color="button" class="action-cancel" @click.stop="cancel">
          {{ $gettext(`Cancel`) }}
        </v-btn>
        <v-btn color="highlight" variant="flat" class="action-confirm" @click.stop="confirm">
          {{ $gettext(`Yes`) }}
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
<script>
export default {
  name: "PConfirmDialog",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    icon: {
      type: String,
      default: "mdi-delete-outline",
    },
  },
  data() {
    return {};
  },
  methods: {
    cancel() {
      this.$emit("cancel");
    },
    confirm() {
      this.$emit("confirm");
    },
  },
};
</script>
