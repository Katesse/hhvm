/*
 *  Copyright (c) 2023-present, Facebook, Inc.
 *  All rights reserved.
 *
 *  This source code is licensed under the BSD-style license found in the
 *  LICENSE file in the root directory of this source tree.
 */

#include <fizz/util/Tracing.h>
#include <folly/tracing/StaticTracepoint.h>

namespace fizz {
extern "C" {
void fizz_probe_secret_available(
    unsigned char* secretData,
    unsigned char secretSize,
    fizz::KeyLogWriter::Label nssLabel,
    unsigned char* clientRandom) {
  FOLLY_SDT(
      fizz,
      fizz_secret_available,
      secretData,
      secretSize,
      nssLabel,
      clientRandom);
}
}

} // namespace fizz
