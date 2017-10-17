import { browser, by, element, protractor } from 'protractor';

export class FriendshipPage {
  sleep() {
    browser.sleep(3000);
  }

  navigateToHome() {
    return browser.get('/home');
  }

  getFollowButton() {
    return element(by.css('.followbtn'));
  }

  getUserFollowing() {
    return element(by.css('.following'));
  }

  getFriendCol() {
    return element(by.css('mat-card'));
  }

  getNavRightButton() {
    return element(by.css('.navRightButton'));
  }

  getLogoutButton() {
    return element(by.css('.logoutBtn'));
  }

}
