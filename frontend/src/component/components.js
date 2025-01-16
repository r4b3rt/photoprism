/*

Copyright (c) 2018 - 2025 PhotoPrism UG. All rights reserved.

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

import PNotify from "notify.vue";
import PScroll from "scroll.vue";
import PNavigation from "navigation.vue";
import PLoadingBar from "loading-bar.vue";
import PViewer from "viewer.vue";

import PAccountAppsDialog from "account/apps.vue";
import PAccountPasscodeDialog from "account/passcode.vue";
import PAccountPasswordDialog from "account/password.vue";

import PAlbumClipboard from "album/clipboard.vue";
import PAlbumToolbar from "album/toolbar.vue";
import PAlbumEditDialog from "album/dialog/edit.vue";
import PAlbumDeleteDialog from "album/dialog/delete.vue";

import PAuthHeader from "auth/header.vue";
import PAuthFooter from "auth/footer.vue";

import PUploadDialog from "dialog/upload.vue";
import PShareDialog from "dialog/share.vue";
import PWebdavDialog from "dialog/webdav.vue";
import PReloadDialog from "dialog/reload.vue";
import PSponsorDialog from "dialog/sponsor.vue";
import PConfirmDialog from "dialog/confirm.vue";

import PFileClipboard from "file/clipboard.vue";
import PFileDeleteDialog from "file/dialog/delete.vue";

import IconLivePhoto from "icon/live-photo.vue";
import IconSponsor from "icon/sponsor.vue";
import IconPrism from "icon/prism.vue";

import PPeopleMergeDialog from "people/dialog/merge.vue";
import PPeopleEditDialog from "people/dialog/edit.vue";

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

import PServiceAddDialog from "service/dialog/add.vue";
import PServiceRemoveDialog from "service/dialog/remove.vue";
import PServiceEditDialog from "service/dialog/edit.vue";
import PServiceUploadDialog from "service/dialog/upload.vue";

import PLabelClipboard from "label/clipboard.vue";
import PLabelDeleteDialog from "label/dialog/delete.vue";
import PLabelEditDialog from "label/dialog/edit.vue";

import PSubjectClipboard from "subject/clipboard.vue";

// Installs the components imported above.
export function install(app) {
  app.component("PNotify", PNotify);
  app.component("PScroll", PScroll);
  app.component("PNavigation", PNavigation);
  app.component("PLoadingBar", PLoadingBar);
  app.component("PViewer", PViewer);

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

  app.component("IconLivePhoto", IconLivePhoto);
  app.component("IconSponsor", IconSponsor);
  app.component("IconPrism", IconPrism);

  app.component("PLabelClipboard", PLabelClipboard);
  app.component("PLabelDeleteDialog", PLabelDeleteDialog);
  app.component("PLabelEditDialog", PLabelEditDialog);

  app.component("PPeopleMergeDialog", PPeopleMergeDialog);
  app.component("PPeopleEditDialog", PPeopleEditDialog);

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

  app.component("PSubjectClipboard", PSubjectClipboard);
}
