import { FriendshipPage } from './addfriend.po';

describe('twitter-app NavBar', () => {
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

    page.navigateToHome();
    page.sleep();
    expect(following.getText()).toEqual('Following: 1');

  });

});
