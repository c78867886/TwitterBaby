import { NavBarPage } from './navbar.po';

describe('twitter-app NavBar', () => {
  let page: NavBarPage;

  beforeEach(() => {
    page = new NavBarPage();
  });

  it('should search a friend', () => {
    let searchInput = page.getNavBarSearchInput();
    
    // Search some one who is not in database
    searchInput.sendKeys('MrRight');
    page.sleep();
    page.pressEnterKey();
    page.sleep();
    
    // Search someone who is in database
    searchInput.sendKeys('MarsLee');
    page.sleep();
    page.pressEnterKey();
    page.sleep();
  });

});
