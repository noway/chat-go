'use strict';

angular.module('myApp.messenger', [])

.component('messenger', {
  templateUrl: 'components/messenger/messenger.html',
  controller: ['$scope', 'profileService', function messengerController($scope, profileService) {
    $scope.profile = profileService;    
  }],
  bindings: {
    // messenger: '='
  }
});
