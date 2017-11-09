import { FriendshipPage } from './addfriend.po';

describe('twitter-app Friendship', () => {
  let page: FriendshipPage;

  beforeEach(() => {
    page = new FriendshipPage();
  });

  it('should add a friend', () => {
    let button = page.getFollowButton();
    let following = page.getUserFollowing();
    expect(button.getText()).toEqual('Follow');

    button.click();
    expect(button.getText()).toEqual('Following');

  });

  it('should have a firend in following list', () => {
    page.navigateToHome();
    page.sleep();
    let following = page.getUserFollowing();
    expect(following.getText()).toEqual('Following: 1');
    following.click();
    page.sleep();
    expect(page.getFriendCol()).toBeTruthy();
  });

  it('should show profile page', () => {
    let rightBtn = page.getNavRightButton();
    rightBtn.click();
    page.sleep();
    let profileBtn = page.getProfileBtn();
    profileBtn.click();
    page.sleep();
    expect(page.getProfileCpnt().isPresent()).toBe(true);
  })

  it('should log out', () => {
    let rightBtn = page.getNavRightButton();
    rightBtn.click();
    page.sleep();
    let logoutBtn = page.getLogoutButton();
    logoutBtn.click();
    page.sleep();
  });

});
