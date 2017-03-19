'use strict';

angular.module('myApp.messenger', [])

.component('messenger', {
  templateUrl: 'components/messenger/messenger.html',

  controller: ['$scope', '$http', '$httpParamSerializer', 'profileService', 'statsService', 
  function messengerController($scope, $http, $httpParamSerializer, profileService, statsService) {

    $scope.profile = profileService;    
    
    $scope.send = event => {
      if (!event.message) {
        return;
      }

      return $http({
        url: '/messages',
        method: 'POST',
        data: $httpParamSerializer(event), 
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded' // Note the appropriate header
        }
      }).then(res => {
        // if (res.data.Code != 0) {
        //   return alert(`Exit code: ${res.data.Code}. Error message: ${res.data.Message}`);
        // }
        $scope.event.message = '';
        statsService.updateSendMessage();
      });

    };

  }]
});
