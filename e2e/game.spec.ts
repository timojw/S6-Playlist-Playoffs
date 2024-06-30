import { test as base, expect } from '@playwright/test';

const test = base.extend({
  page: async ({ browser }, use) => {
    const context = await browser.newContext({
      ignoreHTTPSErrors: true,
    });
    await use(await context.newPage());
    await context.close();
  },
});

test.describe('Game functionality', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('https://4.182.90.170.nip.io', { waitUntil: 'networkidle' });
  });

  test('can search game by ID', async ({ page }) => {
    await page.fill('input[placeholder="Enter game ID"]', '1');
    await page.click('text=Search');
    await expect(page).toHaveURL('https://4.182.90.170.nip.io/game/1');
  });

  test('can add a new game', async ({ page }) => {
    await page.goto('https://4.182.90.170.nip.io/start-game');
    await page.fill('input#name', 'New Game');
    await page.fill('input#deadline', '2023-01-01');
    await page.click('button[type="submit"]');
    // Further assertions to check if the game was added
  });
});
