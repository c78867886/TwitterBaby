import { browser, by, element, protractor } from 'protractor';

export class NavBarPage {
  sleep() {
    browser.sleep(1000);
  }

  navigateToHome() {
    return browser.get('/home');
  }

  //Test for Nav Bar Friend Search
  getNavBarSearchInput() {
    return element(by.css('input[id=input]'));
  }

  pressEnterKey() {
    browser.actions().sendKeys(protractor.Key.ENTER).perform();
  }
}
