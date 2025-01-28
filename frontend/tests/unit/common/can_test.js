import "../fixtures";
import * as can from "common/can";

let chai = require("chai/chai");
let assert = chai.assert;

describe("common/can", () => {
  it("canUseVideo", () => {
    assert.equal(can.useVideo, true);
  });

  it("canUseAvc", () => {
    assert.equal(can.useAvc, true);
  });

  it("canUseOGV", () => {
    assert.equal(can.useOGV, true);
  });

  it("canUseVP8", () => {
    assert.equal(can.useVp8, false);
  });

  it("canUseVP9", () => {
    assert.equal(can.useVp9, true);
  });

  it("canUseAv1", () => {
    assert.equal(can.useAv1, true);
  });

  it("canUseWebM", () => {
    assert.equal(can.useWebM, true);
  });

  it("canUseHevc", () => {
    assert.equal(can.useHevc, false);
  });
});
