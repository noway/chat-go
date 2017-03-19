angular.module('myApp.stateService', [])

.service('stateService', ['$q', '$http', '$rootScope', function ($q, $http, $rootScope) {

	this.fetchState = () => $http.get('/users/state').then(res => res.data);
	this.fetchOnline = () => $http.get('/users/online').then(res => res.data);

	this.updateState = () => (this.getState = this.fetchState()).then(() => $rootScope.$broadcast('updateState'))
	this.updateOnline = () => (this.getOnline = this.fetchOnline()).then(() => $rootScope.$broadcast('updateOnline'))
	
	this.updateState();
	this.updateOnline();
}]);