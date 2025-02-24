/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <iostream>
#include "util_nas5g_pkt.h"

namespace magma5g {

//  API for testing decode registration request
bool decode_registration_request_msg(
    RegistrationRequestMsg* reg_request, const uint8_t* buffer, uint32_t len) {
  bool decode_success = true;
  uint8_t* decode_reg_buffer =
      const_cast<uint8_t*>(reinterpret_cast<const uint8_t*>(buffer));

  if (reg_request->DecodeRegistrationRequestMsg(
          reg_request, decode_reg_buffer, len) < 0) {
    decode_success = false;
  }

  return (decode_success);
}

}  // namespace magma5g
