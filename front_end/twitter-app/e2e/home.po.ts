import { browser, by, element, protractor } from 'protractor';

export class HomePage {
  // Async
  sleep() {
    browser.sleep(3000);
  }

  // Navigate
  navigateToHome() {
    return browser.get('/home');
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

  getNewTweet() {
    return element.all(by.css('.tweetMessage')).first();
  }
}
