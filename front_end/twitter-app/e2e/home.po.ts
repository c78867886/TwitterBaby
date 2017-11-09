import { browser, by, element, protractor } from 'protractor';

export class HomePage {
  // Async
  sleep() {
    browser.sleep(1000);
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

  getCommentBtn() {
    return element.all(by.css('#comment')).first();
  }

  getCommentText() {
    return element(by.css('.expansionpanel2'));
  }

  getCommentTextArea() {
    return element(by.css('.commenttextarea'));
  }

  getCommentSubmit() {
    return element(by.css('.actBtn2'));
  }

  getCommentList() {
    return element.all(by.css('.commentContent')).first();
  }

  getCancelBtn() {
    return element(by.css("#closeBtn"));
  }
}
