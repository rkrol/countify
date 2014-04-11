'use strict';

var accountServices = angular.module('accountServices', ['ngResource']);

accountServices.factory('Account', ['$resource',
  function($resource){
    return $resource('accounts/:accountId', {
      accountId: '@accountId'
    }, {
      saveAndList: {method:'POST', isArray:true}
    });
  }]);
