angular.module('myApp.profileService', ['myApp.stateService'])

.service('profileService', ['$rootScope', 'stateService', function ($rootScope, stateService) {
	
	this.user = '';

    $rootScope.$on('updateState', () => stateService.getState.then(e => {
		this.user = e.Message
	}))

	this.updateUser = () => {
		stateService.updateState();
		stateService.updateOnline();
	}
}]);