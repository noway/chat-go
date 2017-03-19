'use strict';

angular.module('myApp.profile', ['myApp.profileService'])

.component('profile', {
  templateUrl: 'components/profile/profile.html',
  controller: ['$scope', '$http', 'profileService', function profileController($scope, $http, profileService) {
    $scope.profile = profileService;
    
    $scope.login = user => $http.get('/users/login', { params: user }).then(res => {
      if (res.data.Code != 0) {
        return alert(`Exit code: ${res.data.Code}. Error message: ${res.data.Message}`);
      }
      profileService.updateUser();
    });

    $scope.register = user => $http.get('/users/register', { params: user }).then(res => {
      if (res.data.Code != 0) {
        return alert(`Exit code: ${res.data.Code}. Error message: ${res.data.Message}`);
      }
      profileService.updateUser();
    });

    $scope.logout = () => $http.get('/users/logout').then(res => {
      if (res.data.Code != 0) {
        return alert(`Exit code: ${res.data.Code}. Error message: ${res.data.Message}`);
      }
      profileService.updateUser();
    });
    

  }]

});
