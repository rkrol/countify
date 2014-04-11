'use strict';

angular
  .module('countifyApp', [
    'ngCookies',
    'ngResource',
    'ngSanitize',
    'ngRoute',
    'accountServices',
    'accountControllers'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/accounts', {
        templateUrl: 'views/main.html',
        controller: 'AccountListController'
      })
      .when('/accounts/new', {
        templateUrl: 'views/new.html',
        controller: 'AccountListController'
      })
      .when('/account/:accountId', {
        templateUrl: 'views/account.html',
        controller: 'AccountDetailsController'
      })
      .otherwise({
        redirectTo: '/accounts'
      });
  });
