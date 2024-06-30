import { test, expect } from '@playwright/test';

test('can login', async ({ page }) => {
  await page.goto('/');

  await page.getByRole('button', { name: 'Log in' }).click();

  await page
    .getByLabel('Email address*')
    .fill("PLAYWRIGHT_USER@mail.com" || '');
  await page
    .getByLabel('Password*')
    .fill("PLAYWRIGHT_PASSWORD1" || '');
  await page.getByRole('button', { name: 'Continue', exact: true }).click();
  await expect(page).toHaveURL('/');
});

test('can logout', async ({ page }) => {
  await page.goto('/');

  await page.getByRole('button', { name: 'Log in' }).click();

  await page
    .getByLabel('Email address*')
    .fill("PLAYWRIGHT_USER@mail.com" || '');
  await page
    .getByLabel('Password*')
    .fill("PLAYWRIGHT_PASSWORD1" || '');
  await page.getByRole('button', { name: 'Continue', exact: true }).click();
  await expect(page).toHaveURL('/');

  await page.click('text=Log out');
});

test('can see game create page', async ({ page }) => {
  await page.goto('/');

  await page.getByRole('button', { name: 'Log in' }).click();

  await page
    .getByLabel('Email address*')
    .fill("PLAYWRIGHT_USER@mail.com" || '');
  await page
    .getByLabel('Password*')
    .fill("PLAYWRIGHT_PASSWORD1" || '');
  await page.getByRole('button', { name: 'Continue', exact: true }).click();
  await expect(page).toHaveURL('/');
  await page.getByText('Create Game').click();
  await expect(page).toHaveURL('/start-game');
});
