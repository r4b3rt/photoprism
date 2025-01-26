<template>
  <div class="p-tab p-tab-photo-people">
    <div class="pa-2 p-faces">
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
          <v-card :data-id="marker.UID" :class="marker.classes()" class="result not-selectable flex-grow-1">
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
                hide-details
                single-line
                clearable
                persistent-clear
                clear-icon="mdi-eject"
                density="comfortable"
                class="input-name pa-0 ma-0"
                @click:clear="onClearSubject(marker)"
              ></v-text-field>
              <!-- TODO: check property allow-overflow TEST -->
              <v-combobox
                v-else
                v-model:search="marker.Name"
                :items="$config.values.people"
                item-title="Name"
                item-value="Name"
                :disabled="busy"
                return-object
                hide-no-data
                :menu-props="menuProps"
                hide-details
                single-line
                open-on-clear
                focused
                append-icon=""
                prepend-inner-icon="mdi-account-plus"
                density="comfortable"
                class="input-name pa-0 ma-0"
                @blur="onRename(marker)"
                @update:model-value="(person) => onUpdate(marker, person)"
                @keyup.enter.native="onRename(marker)"
              >
              </v-combobox>
            </v-card-actions>
          </v-card>
        </div>
      </div>
    </div>
    <p-confirm-action
      :show="confirm.show"
      icon="mdi-account-plus"
      :icon-size="42"
      :text="confirm.marker?.Name ? $gettext('Add %{name}?', { name: confirm.marker.Name }) : $gettext('Add person?')"
      @close="onRenameCancelled"
      @confirm="onRenameConfirmed"
    ></p-confirm-action>
  </div>
</template>

<script>
import Marker from "model/marker";
import PConfirmAction from "component/confirm/action.vue";
import { reactive } from "vue";

export default {
  name: "PTabPhotoPeople",
  components: { PConfirmAction },
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
      confirm: {
        show: false,
        marker: new Marker(),
        text: this.$gettext("Add person?"),
      },
      menuProps: {
        closeOnClick: false,
        closeOnContentClick: true,
        openOnClick: false,
        density: "compact",
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
    onUpdate(marker, person) {
      if (typeof person === "object" && marker?.UID && person?.UID && person?.Name) {
        marker.Name = person.Name;
        marker.SubjUID = person.UID;
        this.rename(marker);
      }

      return true;
    },
    onRename(marker) {
      if (this.busy || !marker) {
        return;
      }

      const name = marker?.Name;

      if (!name) {
        this.onRenameCancelled();
        return;
      }

      this.confirm.marker = marker;

      const people = this.$config.values?.people;

      if (people) {
        const found = people.find((person) => person.Name.localeCompare(name, "en", { sensitivity: "base" }) === 0);
        if (found) {
          marker.Name = found.Name;
          marker.SubjUID = found.UID;
          this.rename(marker);
          return;
        }
      }

      marker.Name = name;
      marker.SubjUID = "";
      this.confirm.show = true;
    },
    onRenameConfirmed() {
      if (!this.confirm?.marker?.Name) {
        return;
      }

      this.rename(this.confirm.marker);
    },
    onRenameCancelled() {
      if (this.confirm?.marker?.Name && this.confirm?.marker?.originalValue) {
        // Revert name change.
        this.confirm.marker.Name = this.confirm.marker.originalValue("Name");
      }

      this.$nextTick(() => {
        this.confirm.marker = reactive(new Marker());
        this.confirm.show = false;
      });
    },
    rename(marker) {
      if (this.busy || !marker) return;

      this.busy = true;
      this.$notify.blockUI();

      marker.rename().finally(() => {
        this.$notify.unblockUI();
        this.busy = false;
        this.confirm.marker = reactive(new Marker());
        this.confirm.show = false;
      });
    },
  },
};
</script>
