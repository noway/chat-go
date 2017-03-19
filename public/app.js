'use strict';

// Declare app level module which depends on views, and components
angular.module('myApp', [
  'ngRoute',
  'myApp.page',
  'myApp.profile',
  'myApp.paginator',
  'myApp.messenger',
  'myApp.stats',
])
.config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
	
  $locationProvider.hashPrefix('');
  $routeProvider.otherwise({ redirectTo: 'page/0' });
  
}])