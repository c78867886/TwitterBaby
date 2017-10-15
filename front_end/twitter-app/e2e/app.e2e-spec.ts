import { AppPage } from './app.po';

describe('twitter-app App', () => {
  let page: AppPage;

  beforeEach(() => {
    page = new AppPage();
  });

  it('should display signup page', () => {
    page.navigateToSign();
    expect(page.getParagraphText()).toEqual('Sign Up Twitter Baby Today!');
  });
  
  it('should display login page and log in', () => {
    page.navigateToLogin();
    expect(page.getParagraphText()).toEqual('Welcome to Twitter Baby!');
    let usernameInput = page.loginUsername();
    let passwordInput = page.loginPassword();
    let loginSubmitButton = page.loginSubmit();

    usernameInput.sendKeys('hojason117@gmail.com');
    expect(usernameInput.getAttribute('value')).toEqual('hojason117@gmail.com');

    passwordInput.sendKeys('test1');
    expect(passwordInput.getAttribute('value')).toEqual('test1');
    loginSubmitButton.click();
    page.sleep();
  });

  it('should display home page', () => {
    //page.navigateToHome();
    expect(page.getNavBarBrand()).toEqual('BabyTwitter');
  });

  it('should post a new tweet', () => {
    page.getPostArea().click();
    expect(page.getPostSubmit().getAttribute('disabled')).toBe('true');

    page.getPostInputArea().sendKeys('Test Message for E2E test!');
    expect(page.getPostInputArea().getAttribute('value')).toEqual('Test Message for E2E test!');
    expect(page.getPostSubmit().getAttribute('disabled')).toBe(null);

    page.getPostSubmit().click();
    page.sleep();
    
  });
  
  it('should search a friend', () => {
    page.getNavBarSearchInput().sendKeys('MarsLee');
    page.pressEnterKey();
    page.sleep();
  });

});
