'use strict';

angular.module('myApp.stats', ['myApp.profileService', 'myApp.statsService'])

.component('stats', {
  templateUrl: 'components/stats/stats.html',
  controller: ['$scope', 'statsService', function profileController($scope, statsService) {

  	$scope.stats = statsService;
    
    // $scope.usersOnline = [];
    // $scope.memoryRss = statsService.usersOnline
  }],

  // bindings: {
  //   // stats: '=',
  //   // memoryRss: '<',
  //   my: '<',
  // }
});
