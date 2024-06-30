import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './e2e',  // Ensure this points to your e2e test directory
  use: {
    baseURL: 'https://4.182.90.170.nip.io/', // Set the base URL for all tests
    ignoreHTTPSErrors: true, // This will ignore HTTPS errors, useful for self-signed certificates
    browserName: 'chromium', // if you prefer to default to a specific browser
    headless: false, // Set to false if you want to see the browser while testing
    viewport: { width: 1280, height: 720 } // Set a default viewport size
  },
  projects: [
    {
      name: 'Desktop Chromium',
      use: {
        browserName: 'chromium',
        // Any other specific settings for Chromium
      },
    },
    {
      name: 'Desktop Firefox',
      use: {
        browserName: 'firefox',
        // Any other specific settings for Firefox
      },
    },
    {
      name: 'Desktop WebKit',
      use: {
        browserName: 'webkit',
        // Any other specific settings for WebKit
      },
    },
  ],
  // Additional global configurations can go here
});
