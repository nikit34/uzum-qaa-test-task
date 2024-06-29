*** Settings ***
Library    SeleniumLibrary

*** Variables ***
${URL}    https://release-b2c-stand.dev.internal.daymarket.uz/ru
${BROWSER}    chrome

*** Test Cases ***
Open Homepage and Verify Title
    [Documentation]    Открытие главной страницы и проверка заголовка
    Open Browser    ${URL}    ${BROWSER}
    Maximize Browser Window
    Title Should Be    Uzum Market — интернет-магазин
    Close Browser

Search for a Product
    [Documentation]    Поиск товара и проверка результатов
    Open Browser    ${URL}    ${BROWSER}
    Maximize Browser Window
    Wait Until Page Contains Element    //*[@id="search-form"]/div/div[1]/input
    Input Text    //*[@id="search-form"]/div/div[1]/input    ноутбук
    Click Button    //*[@id="search-form"]/button
    Wait Until Page Contains Element    //*[@id="category-products"]
    Element Should Be Visible    //*[@id="category-products"]/div[1]/div
    Close Browser

Navigate to Login Page
    [Documentation]    Переход на модальное окно логина
    Open Browser    ${URL}    ${BROWSER}
    Maximize Browser Window
    Wait Until Page Contains Element    //div[@data-test-id="button__auth"]
    Click Element    //div[@data-test-id="button__auth"]
    Wait Until Element Is Visible    //div[@class="sign-in-phone"]
    Close Browser

*** Keywords ***
Fail And Capture Screenshot
    [Arguments]    ${message}
    Capture Page Screenshot
    Fail    ${message}
    Remove File    screenshot.png

*** Test Teardown ***
    Run Keyword If Test Failed    Fail And Capture Screenshot    Test failed
