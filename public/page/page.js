'use strict';

angular.module('myApp.page', ['ngRoute'])

.config(['$routeProvider', function ($routeProvider) {
  $routeProvider.when('/page/:page', {
    templateUrl: 'page/page.html',
    controller: 'PageCtrl'
  })
}])

.controller('PageCtrl', ['$scope', '$routeParams', '$http', function ($scope, $routeParams, $http) {  
  $scope.last = 0;
  $scope.events = [];
  $scope.shownEvents = {};
  $scope.user = '';
  
  $scope.processMessages = events => {

    events.map(e => {
      if (!$scope.shownEvents[e.X]) {
        $scope.events.push(e);
        $scope.shownEvents[e.X] = e;
        $scope.last = e.X >= $scope.last ? e.X : $scope.last; 
      }
    });
    
    if ($routeParams.page == 0) {
      $scope.pollMessages();
    }

  };

  $scope.getMessages = () => $http.get('/load?page=' +  $routeParams.page).then(res => res.data).then($scope.processMessages);
  $scope.pollMessages = () => $http.get('/messages?last=' + $scope.last).then(res => res.data).then($scope.processMessages);

  $scope.getState = () => $http.get('/users/state').then(e => {
    $scope.memoryRss = e.RSS;
    $scope.user = e.Message;
  });

  $scope.date = event => (new Date(event.T*1000)).toLocaleString();
  $scope.name = event => event.N == '' ? 'Аноним' : event.N;
  $scope.isMine = event => event.N == $scope.user && event.L;
  $scope.isMention = event => event.M.indexOf($scope.user) >= 0 && $scope.user.length;
  

  $scope.getState();
  $scope.getMessages();


}]);