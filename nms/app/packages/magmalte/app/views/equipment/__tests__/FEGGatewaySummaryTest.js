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
 *
 * @flow strict-local
 * @format
 */

import 'jest-dom/extend-expect';
import FEGGatewayContext from '../../../components/context/FEGGatewayContext';
import FEGGatewaySummary from '../FEGGatewaySummary';
import MuiStylesThemeProvider from '@material-ui/styles/ThemeProvider';
import React from 'react';
import defaultTheme from '@fbcnms/ui/theme/default';
import {MemoryRouter, Route} from 'react-router-dom';
import {MuiThemeProvider} from '@material-ui/core/styles';
import {cleanup, render, wait} from '@testing-library/react';
import type {federation_gateway} from '@fbcnms/magma-api';

jest.mock('axios');
jest.mock('@fbcnms/magma-api');
jest.mock('@fbcnms/ui/hooks/useSnackbar');
afterEach(cleanup);

const mockHardwareId = 'c9439d30-61ef-46c7-93f2-e01fc131244d';

const mockKeyType = 'SOFTWARE_ECDSA_SHA256';

const mockGw0: federation_gateway = {
  id: 'test_feg_gw0',
  name: 'test_gateway',
  description: 'hello I am a federated gateway',
  tier: 'default',
  device: {
    key: {key: '', key_type: mockKeyType},
    hardware_id: mockHardwareId,
  },
  magmad: {
    autoupgrade_enabled: true,
    autoupgrade_poll_interval: 300,
    checkin_interval: 60,
    checkin_timeout: 100,
    tier: 'tier2',
  },
  federation: {
    aaa_server: {},
    eap_aka: {
      plmn_ids: [],
    },
    gx: {},
    gy: {},
    health: {
      health_services: [],
    },
    hss: {},
    s6a: {},
    served_network_ids: [],
    swx: {
      hlr_plmn_ids: [],
      server: {
        protocol: 'tcp',
      },
      servers: [],
    },
    csfb: {},
  },
  status: {
    checkin_time: 0,
    meta: {
      gps_latitude: '0',
      gps_longitude: '0',
      gps_connected: '0',
      enodeb_connected: '0',
      mme_connected: '0',
    },
    platform_info: {
      packages: [{version: '1.2'}],
    },
  },
};

const mockGw1: federation_gateway = {
  ...mockGw0,
  id: 'test_gw1',
  name: 'test_gateway1',
};

const fegGateways = {
  [mockGw0.id]: mockGw0,
  [mockGw1.id]: mockGw1,
};

const fegGatewaysHealth = {
  [mockGw0.id]: 'HEALTHY',
  [mockGw1.id]: 'UNHEALTHY',
};

describe('<FEGEquipmentGateway />', () => {
  const Wrapper = () => (
    <MemoryRouter
      initialEntries={[
        '/nms/mynetwork/equipment/overview/gateway/test_feg_gw0/overview',
      ]}
      initialIndex={0}>
      <MuiThemeProvider theme={defaultTheme}>
        <MuiStylesThemeProvider theme={defaultTheme}>
          <FEGGatewayContext.Provider
            value={{
              state: fegGateways,
              setState: async _ => {},
              health: fegGatewaysHealth,
              activeFegGatewayId: mockGw0.id,
            }}>
            <Route
              path="/nms/:networkId/equipment/overview/gateway/:gatewayId/overview"
              render={props => (
                <FEGGatewaySummary {...props} gwInfo={mockGw0} />
              )}
            />
          </FEGGatewayContext.Provider>
        </MuiStylesThemeProvider>
      </MuiThemeProvider>
    </MemoryRouter>
  );

  it('renders federation gateway summary correctly', async () => {
    const {getByTestId} = render(<Wrapper />);
    await wait();
    // verify gateway information
    expect(getByTestId('Name')).toHaveTextContent('test_gateway');
    expect(getByTestId('Gateway ID')).toHaveTextContent('test_feg_gw0');
    expect(getByTestId('Hardware UUID')).toHaveTextContent(
      'c9439d30-61ef-46c7-93f2-e01fc131244d',
    );
    expect(getByTestId('Version')).toHaveTextContent('1.2');
  });
});
