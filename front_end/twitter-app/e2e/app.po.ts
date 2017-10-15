import { browser, by, element, protractor } from 'protractor';

export class AppPage {
  // Async
  sleep() {
    browser.sleep(3000);
  }

  // Navigate
  navigateToLogin() {
    browser.driver.manage().window().setSize(1280, 1024);
    return browser.get('/login');
  }

  navigateToSign() {
    return browser.get('/signup');
  }

  navigateToHome() {
    return browser.get('/home');
  }

  getParagraphText() {
    return element(by.css('h1')).getText();
  }


  // Test for Login Page
  loginUsername() {
    return element(by.css('input[name=username]'));
  }

  loginPassword() {
    return element(by.css('input[name=password]'));
  }

  loginSubmit() {
    browser.ignoreSynchronization = true;
    return element(by.css('form .btn'));
  }

  //Test for Home Page
  getNavBarBrand() {
    return element(by.css('.brand')).getText();
  }

  getPostArea() {
    return element(by.css('.expansionpanel'));
  }

  getPostInputArea() {
    return element(by.css('textarea'));
  }

  getPostSubmit() {
    return element(by.css('.submitbtn'));
  }

  //Test for Nav Bar Friend Search
  getNavBarSearchInput() {
    return element(by.css('input[id=input]'));
  }

  pressEnterKey() {
    browser.actions().sendKeys(protractor.Key.ENTER).perform();
  }
}
