import { LoginPage } from './login.po';

describe('twitter-app login', () => {
  let page: LoginPage;
  let username: string = "test5";
  beforeEach(() => {
    page = new LoginPage();
  });
  
  it('should register for an existing user', () => {
    page.navigateToSign();
    let signUsername = page.signUsername();
    let signEmail = page.signEmail();
    let signFirstName = page.signFirstname();
    let signLastName = page.signLastname();
    let signPassword = page.signPassword();

    signUsername.sendKeys("jasonhe");
    signEmail.sendKeys("hexing_h@hotmail.com");
    signFirstName.sendKeys("Jason");
    signLastName.sendKeys("He");
    signPassword.sendKeys("testpassword");

    expect(signPassword.getAttribute('value')).toEqual("testpassword");
    expect(signFirstName.getAttribute('value')).toEqual("Jason");
    expect(signLastName.getAttribute('value')).toEqual("He");
    page.sleep();
    page.signSubmitBtn().click();
    page.sleep();
    expect(page.alertDiv().getText()).toEqual('Something is wrong, please sign up again.');
    page.sleep();
  });

  it('should display signup page', () => {
    page.navigateToSign();
    let signUsername = page.signUsername();
    let signEmail = page.signEmail();
    let signFirstName = page.signFirstname();
    let signLastName = page.signLastname();
    let signPassword = page.signPassword();

    signUsername.sendKeys(username);
    signEmail.sendKeys(username + "@gmail.com");
    signFirstName.sendKeys("Jason");
    signLastName.sendKeys("He");
    signPassword.sendKeys("testpassword");

    expect(signUsername.getAttribute('value')).toEqual(username);
    expect(signEmail.getAttribute('value')).toEqual(username + "@gmail.com");
    expect(signPassword.getAttribute('value')).toEqual("testpassword");
    expect(signFirstName.getAttribute('value')).toEqual("Jason");
    expect(signLastName.getAttribute('value')).toEqual("He");
    page.sleep();
    page.signSubmitBtn().click();
    page.sleep();
    page.sleep();
    page.sleep();
  });

  it('should display login page and log in', () => {
    //page.navigateToLogin();
    expect(page.getParagraphText()).toEqual('Welcome to Twitter Baby!');
    let usernameInput = page.loginUsername();
    let passwordInput = page.loginPassword();
    let loginSubmitButton = page.loginSubmit();

    usernameInput.sendKeys(username + '@gmail.com');
    expect(usernameInput.getAttribute('value')).toEqual(username + '@gmail.com');

    passwordInput.sendKeys('testpassword');
    expect(passwordInput.getAttribute('value')).toEqual('testpassword');
    page.sleep();
    loginSubmitButton.click();
    page.sleep();
  });

});
