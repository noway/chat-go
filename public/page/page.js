'use strict';

angular.module('myApp.page', ['ngRoute'])

.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/page/:page', {
    templateUrl: 'page/page.html',
    controller: 'PageCtrl'
  })
}])

.controller('PageCtrl', ['$scope', '$routeParams', '$http', function ($scope, $routeParams, $http) {
	// var page = 0;
	// page = parseInt(location.hash.substr(1), 10) || 0;
	
	$scope.last = 0;
	$scope.events = [];
	$scope.user = '';

	$scope.getMessages = function () {
		$http.get('/load?page=' +  $routeParams.page).then(res => {
			angular.copy(res.data, $scope.events);
			res.data.map(e => $scope.last = e.X);
			
			if ($routeParams.page == 0) {
				$scope.pollMessages();
			}
		});
	};

	$scope.getState = function () {
		$http.get('/users/state').then(e => {
			
			$scope.memoryRss = e.RSS;
			// if (e.Message !=  ''){
			// 	$scope.user = false;
			// 	$('#logout-form').show()
			// 	$('#name').hide()
			// } else {
			// 	$scope.user = e.Message;
			// 	$('#login-form').show()
			// }
			$scope.user = e.Message;
			// getMessages();
		});
	};
	

	$scope.date = function(event) {
		return (new Date(event.T*1000)).toLocaleString() 
	};
	
	$scope.name = function(event) {
		return event.N == '' ? 'Аноним' : event.N;	
	};
	
	$scope.isMine = function(event) {
		return event.N == $scope.user && event.L;
	};
	$scope.isMention = function(event) {
		return event.M.indexOf($scope.user) >= 0 && $scope.user.length;
	};
	
	$scope.pollMessages = function () {
		$http.get('/messages?last=' + $scope.last).then(res => {
			angular.copy($scope.events, res.data);
			$scope.pollMessages();
		});
	};
	$scope.getState();
	$scope.getMessages();


}]);