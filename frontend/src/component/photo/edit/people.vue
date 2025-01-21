<template>
  <div class="p-tab p-tab-photo-people">
    <v-container grid-list-xs fluid class="pa-2 p-faces">
      <v-alert
        v-if="markers.length === 0"
        color="surface-variant"
        icon="mdi-lightbulb-outline"
        class="no-results ma-2 opacity-70"
        variant="outlined"
      >
        <div class="font-weight-bold">
          {{ $gettext(`No people found`) }}
        </div>
        <div class="mt-2">
          {{ $gettext(`You may rescan your library to find additional faces.`) }}
          {{ $gettext(`Recognition starts after indexing has been completed.`) }}
        </div>
      </v-alert>
      <div v-else class="v-row search-results face-results cards-view d-flex">
        <div v-for="marker in markers" :key="marker.UID" class="v-col-12 v-col-sm-6 v-col-md-4 v-col-lg-3 d-flex">
          <v-card :data-id="marker.UID" style="user-select: none" :class="marker.classes()" class="result flex-grow-1">
            <v-img :src="marker.thumbnailUrl('tile_320')" :transition="false" aspect-ratio="1" class="card">
              <v-btn
                v-if="!marker.SubjUID && !marker.Invalid"
                :ripple="false"
                class="input-reject"
                icon
                variant="text"
                density="comfortable"
                position="absolute"
                :title="$gettext('Remove')"
                @click.stop.prevent="onReject(marker)"
              >
                <v-icon class="action-reject">mdi-close</v-icon>
              </v-btn>
            </v-img>
            <v-card-actions class="meta pa-0">
              <v-btn
                v-if="marker.Invalid"
                :disabled="busy"
                size="large"
                variant="flat"
                block
                :rounded="false"
                class="action-undo text-center"
                :title="$gettext('Undo')"
                @click.stop="onApprove(marker)"
              >
                <v-icon>mdi-undo</v-icon>
              </v-btn>
              <v-text-field
                v-else-if="marker.SubjUID"
                v-model="marker.Name"
                :rules="[textRule]"
                :disabled="busy"
                :readonly="true"
                autocomplete="off"
                autocorrect="off"
                class="input-name pa-0 ma-0"
                hide-details
                single-line
                clearable
                persistent-clear
                clear-icon="mdi-eject"
                @click:clear="onClearSubject(marker)"
                @change="onRename(marker)"
                @keyup.enter="onRename(marker)"
              ></v-text-field>
              <!-- TODO: check property allow-overflow TEST -->
              <v-combobox
                v-else
                v-model="marker.Name"
                style="z-index: 250"
                :items="$config.values.people"
                item-title="Name"
                item-value="Name"
                :disabled="busy"
                :return-object="false"
                :menu-props="menuProps"
                :hint="$gettext('Name')"
                hide-details
                single-line
                open-on-clear
                hide-no-data
                append-icon=""
                prepend-inner-icon="mdi-account-plus"
                autocomplete="off"
                class="input-name pa-0 ma-0"
                @blur="onRename(marker)"
                @keyup.enter.native="onRename(marker)"
              >
              </v-combobox>
            </v-card-actions>
          </v-card>
        </div>
      </div>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "PTabPhotoPeople",
  props: {
    model: {
      type: Object,
      default: () => {},
    },
    uid: {
      type: String,
      default: "",
    },
  },
  data() {
    return {
      busy: false,
      markers: this.model.getMarkers(true),
      imageUrl: this.model.thumbnailUrl("fit_720"),
      disabled: !this.$config.feature("edit"),
      config: this.$config.values,
      readonly: this.$config.get("readonly"),
      menuProps: {
        closeOnClick: false,
        closeOnContentClick: true,
        openOnClick: false,
        maxHeight: 300,
      },
      textRule: (v) => {
        if (!v || !v.length) {
          return this.$gettext("Name");
        }

        return v.length <= this.$config.get("clip") || this.$gettext("Name too long");
      },
    };
  },
  watch: {
    model: function () {
      this.refresh();
    },
  },
  methods: {
    refresh() {
      this.markers = this.model.getMarkers(true);
      this.imageUrl = this.model.thumbnailUrl("fit_720");
    },
    onReject(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.reject().finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onApprove(marker) {
      if (this.busy || !marker) return;

      this.busy = true;

      marker.approve().finally(() => (this.busy = false));
    },
    onClearSubject(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.clearSubject(marker).finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
    onRename(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.rename().finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
      });
    },
  },
};
</script>
