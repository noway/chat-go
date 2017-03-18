'use strict';

angular.module('myApp.profile', [])

.component('profile', {
  templateUrl: 'components/profile/profile.html',
  controller: function ProfileController() {
    
  },
  bindings: {
    // profile: '=',
    my: '<'
  }
});
