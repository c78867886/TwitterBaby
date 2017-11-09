import { NavBarPage } from './navbar.po';

describe('twitter-app NavBar', () => {
  let page: NavBarPage;

  beforeEach(() => {
    page = new NavBarPage();
  });

  it('should search a friend', () => {
    let searchInput = page.getNavBarSearchInput();
    
    // Search someone who is not in database
    searchInput.sendKeys('MrRight');
    page.pressEnterKey();
    page.sleep();
    
    // Search someone who is in database
    searchInput.sendKeys('MarsLee');
    page.pressEnterKey();
    page.sleep();
  });

});
