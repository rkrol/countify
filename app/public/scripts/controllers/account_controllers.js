'use strict';

var accountControllers = angular.module('accountControllers', []);

accountControllers.controller('AccountListController', function ($scope, $location, Account) {

    $scope.accounts = Account.query({}, function() {}, function() {
        console.log('error retrieving accounts');
      });

    $scope.newAccount = function() {
      var account = {
        name: $scope.newAccountName,
        creationUserId: '027ebc71-abc6-4aba-b196-849146443614',
        creationDate: '2014-04-02T19:20:30+01:00'
      };
      $scope.accounts = Account.saveAndList(account, function(){
        $location.path('/');
      }, function(){
        console.log('error creation account');
      });
    };
  });

accountControllers.controller('AccountDetailsController', function($scope, $routeParams, Account) {
    console.log($routeParams.accountId);
    $scope.account = Account.get({accountId: $routeParams.accountId});
  });



