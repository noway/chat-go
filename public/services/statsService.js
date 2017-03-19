angular.module('myApp.statsService', ['myApp.stateService'])

.service('statsService', ['$rootScope', '$q', '$http', 'stateService', function ($rootScope, $q, $http, stateService) {

	this.memoryRss = NaN;
	this.usersOnline = [];
	this.guestsCount = NaN;

    $rootScope.$on('updateState', () => stateService.getState.then(e => {
		this.memoryRss = e.RSS;
	}));

    $rootScope.$on('updateOnline', () => stateService.getOnline.then(data => {
		this.usersOnline = data.filter(val => (val || {}).N);
		this.guestsCount = data.reduce((acc,  val) => (val || {}).N == '' ? acc + 1 : acc, 0);
	}));
}]);