import { test } from "node:test";
import assert from "node:assert/strict";
import { greet } from "./fixture.ts";

await test("greet returns a hello string", () => {
  assert.equal(greet("town"), "hello town");
});
