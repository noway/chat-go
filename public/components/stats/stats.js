'use strict';

angular.module('myApp.stats', [])

.component('stats', {
  templateUrl: 'components/stats/stats.html',
  controller: ['$scope', '$element', '$attrs', '$http', function ProfileController($scope, $element, $attrs, $http) {

  // 	$scope.guestsOnline = NaN;
  // 	$scope.usersOnline = NaN;

  // // 	$http.get('/users/state').then(data => {
		// // $scope.usersOnline = data;
		// // $scope.guestsOnline = data.reduce((acc,  val) => {
		// // 	acc = (val || {}).N == '' ? acc + 1 : acc;			
		// // });
  // // 	});
  	
  //   $scope.isMe = event => {
  //   	return event.N == this.my.nickname && event.L
  //   };
    
  //   $scope.isToMe = event => {
  //   	return this.my.nickname.length && event.M.indexOf(this.my.nickname) >= 0;
  //   };
  }],

  bindings: {
    // stats: '=',
    memoryRss: '<',
    my: '<',
  }
});
