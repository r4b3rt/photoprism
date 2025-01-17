// Global site components.
import PNotify from "notify.vue";
import PScroll from "scroll.vue";
import PNavigation from "navigation.vue";
import PLoadingBar from "loading-bar.vue";
import PViewer from "viewer.vue";

// Icons.
import IconLivePhoto from "icon/live-photo.vue";
import IconSponsor from "icon/sponsor.vue";
import IconPrism from "icon/prism.vue";

// User account management.
import PAccountAppsDialog from "account/apps.vue";
import PAccountPasscodeDialog from "account/passcode.vue";
import PAccountPasswordDialog from "account/password.vue";

// Albums.
import PAlbumClipboard from "album/clipboard.vue";
import PAlbumToolbar from "album/toolbar.vue";
import PAlbumEditDialog from "album/dialog/edit.vue";
import PAlbumDeleteDialog from "album/dialog/delete.vue";

// Login.
import PAuthHeader from "auth/header.vue";
import PAuthFooter from "auth/footer.vue";

// Dialogs.
import PUploadDialog from "dialog/upload.vue";
import PShareDialog from "dialog/share.vue";
import PWebdavDialog from "dialog/webdav.vue";
import PReloadDialog from "dialog/reload.vue";
import PSponsorDialog from "dialog/sponsor.vue";
import PConfirmDialog from "dialog/confirm.vue";

// Originals.
import PFileClipboard from "file/clipboard.vue";
import PFileDeleteDialog from "file/dialog/delete.vue";

// Labels.
import PLabelClipboard from "label/clipboard.vue";
import PLabelDeleteDialog from "label/dialog/delete.vue";
import PLabelEditDialog from "label/dialog/edit.vue";

// People.
import PPeopleMergeDialog from "people/dialog/merge.vue";
import PPeopleEditDialog from "people/dialog/edit.vue";
import PPeopleClipboard from "component/people/clipboard.vue";

// Photos.
import PPhotoCards from "photo/view/cards.vue";
import PPhotoMosaic from "photo/view/mosaic.vue";
import PPhotoList from "photo/view/list.vue";
import PPhotoClipboard from "photo/clipboard.vue";
import PPhotoToolbar from "photo/toolbar.vue";
import PPhotoPreview from "photo/preview.vue";
import PPhotoArchiveDialog from "photo/dialog/archive.vue";
import PPhotoAlbumDialog from "photo/dialog/album.vue";
import PPhotoEditDialog from "photo/dialog/edit.vue";
import PPhotoDeleteDialog from "photo/dialog/delete.vue";

// Settings > Services.
import PServiceAddDialog from "service/dialog/add.vue";
import PServiceRemoveDialog from "service/dialog/remove.vue";
import PServiceEditDialog from "service/dialog/edit.vue";
import PServiceUploadDialog from "service/dialog/upload.vue";

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

  app.component("PUploadDialog", PUploadDialog);
  app.component("PShareDialog", PShareDialog);
  app.component("PWebdavDialog", PWebdavDialog);
  app.component("PReloadDialog", PReloadDialog);
  app.component("PSponsorDialog", PSponsorDialog);
  app.component("PConfirmDialog", PConfirmDialog);

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
