// Global site components.
import PNotify from "component/notify.vue";
import PScroll from "component/scroll.vue";
import PNavigation from "component/navigation.vue";
import PLoadingBar from "component/loading-bar.vue";
import PViewer from "component/viewer.vue";

// Icons.
import IconLivePhoto from "component/icon/live-photo.vue";
import IconSponsor from "component/icon/sponsor.vue";
import IconPrism from "component/icon/prism.vue";

// User account management.
import PAccountAppsDialog from "component/account/apps.vue";
import PAccountPasscodeDialog from "component/account/passcode.vue";
import PAccountPasswordDialog from "component/account/password.vue";

// Albums.
import PAlbumClipboard from "component/album/clipboard.vue";
import PAlbumToolbar from "component/album/toolbar.vue";
import PAlbumEditDialog from "component/album/dialog/edit.vue";
import PAlbumDeleteDialog from "component/album/dialog/delete.vue";

// Login.
import PAuthHeader from "component/auth/header.vue";
import PAuthFooter from "component/auth/footer.vue";

// Dialogs.
import PDialogUpload from "component/dialog/upload.vue";
import PDialogShare from "component/dialog/share.vue";
import PDialogWebdav from "component/dialog/webdav.vue";
import PDialogReload from "component/dialog/reload.vue";
import PDialogSponsor from "component/dialog/sponsor.vue";
import PDialogConfirm from "component/dialog/confirm.vue";

// Originals.
import PFileClipboard from "component/file/clipboard.vue";
import PFileDeleteDialog from "component/file/dialog/delete.vue";

// Labels.
import PLabelClipboard from "component/label/clipboard.vue";
import PLabelDeleteDialog from "component/label/dialog/delete.vue";
import PLabelEditDialog from "component/label/dialog/edit.vue";

// People.
import PPeopleMergeDialog from "component/people/dialog/merge.vue";
import PPeopleEditDialog from "component/people/dialog/edit.vue";
import PPeopleClipboard from "component/people/clipboard.vue";

// Photos.
import PPhotoCards from "component/photo/view/cards.vue";
import PPhotoMosaic from "component/photo/view/mosaic.vue";
import PPhotoList from "component/photo/view/list.vue";
import PPhotoClipboard from "component/photo/clipboard.vue";
import PPhotoToolbar from "component/photo/toolbar.vue";
import PPhotoPreview from "component/photo/preview.vue";
import PPhotoArchiveDialog from "component/photo/archive/dialog.vue";
import PPhotoAlbumDialog from "component/photo/album/dialog.vue";
import PPhotoEditDialog from "component/photo/edit/dialog.vue";
import PPhotoDeleteDialog from "component/photo/delete/dialog.vue";

// Settings > Services.
import PServiceAddDialog from "component/service/add.vue";
import PServiceRemoveDialog from "component/service/remove.vue";
import PServiceEditDialog from "component/service/edit.vue";
import PServiceUploadDialog from "component/service/upload.vue";

// Installs the components imported above.
export function install(app) {
  app.component("PNotify", PNotify);
  app.component("PScroll", PScroll);
  app.component("PNavigation", PNavigation);
  app.component("PLoadingBar", PLoadingBar);
  app.component("PViewer", PViewer);

  app.component("IconLivePhoto", IconLivePhoto);
  app.component("IconSponsor", IconSponsor);
  app.component("IconPrism", IconPrism);

  app.component("PAccountAppsDialog", PAccountAppsDialog);
  app.component("PAccountPasscodeDialog", PAccountPasscodeDialog);
  app.component("PAccountPasswordDialog", PAccountPasswordDialog);

  app.component("PAlbumClipboard", PAlbumClipboard);
  app.component("PAlbumToolbar", PAlbumToolbar);
  app.component("PAlbumEditDialog", PAlbumEditDialog);
  app.component("PAlbumDeleteDialog", PAlbumDeleteDialog);

  app.component("PAuthHeader", PAuthHeader);
  app.component("PAuthFooter", PAuthFooter);

  app.component("PDialogUpload", PDialogUpload);
  app.component("PDialogShare", PDialogShare);
  app.component("PDialogWebdav", PDialogWebdav);
  app.component("PDialogReload", PDialogReload);
  app.component("PDialogSponsor", PDialogSponsor);
  app.component("PDialogConfirm", PDialogConfirm);

  app.component("PFileClipboard", PFileClipboard);
  app.component("PFileDeleteDialog", PFileDeleteDialog);

  app.component("PLabelClipboard", PLabelClipboard);
  app.component("PLabelDeleteDialog", PLabelDeleteDialog);
  app.component("PLabelEditDialog", PLabelEditDialog);

  app.component("PPeopleMergeDialog", PPeopleMergeDialog);
  app.component("PPeopleEditDialog", PPeopleEditDialog);
  app.component("PPeopleClipboard", PPeopleClipboard);

  app.component("PPhotoCards", PPhotoCards);
  app.component("PPhotoMosaic", PPhotoMosaic);
  app.component("PPhotoList", PPhotoList);
  app.component("PPhotoClipboard", PPhotoClipboard);
  app.component("PPhotoToolbar", PPhotoToolbar);
  app.component("PPhotoPreview", PPhotoPreview);
  app.component("PPhotoArchiveDialog", PPhotoArchiveDialog);
  app.component("PPhotoAlbumDialog", PPhotoAlbumDialog);
  app.component("PPhotoEditDialog", PPhotoEditDialog);
  app.component("PPhotoDeleteDialog", PPhotoDeleteDialog);

  app.component("PServiceAddDialog", PServiceAddDialog);
  app.component("PServiceRemoveDialog", PServiceRemoveDialog);
  app.component("PServiceEditDialog", PServiceEditDialog);
  app.component("PServiceUploadDialog", PServiceUploadDialog);
}
