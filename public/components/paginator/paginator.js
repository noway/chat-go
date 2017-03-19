'use strict';

angular.module('myApp.paginator', [])

.component('paginator', {
  templateUrl: 'components/paginator/paginator.html',
  controller: ['$routeParams', '$http', '$scope', function PaginatorController($routeParams, $http, $scope) {

    $scope.prev = NaN;
    $scope.next = NaN;

    if ($routeParams.page == 0) {

      $http.get('/page').then(res => {
        if (res.data.PrevPage == $routeParams.page) {
          return;
        }
        // $('#prev-page').html('<a href="#'+res.data.PrevPage+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + res.data.PrevPage+'</a>');

        $scope.prev = res.data.PrevPage;
      })
    } else {
      $http.get('/page').then(res => {
        if ($routeParams.page > 1) {
          $scope.prev = $routeParams.page - 1;
          // $('#prev-page').html('<a href="#'+($routeParams.page-1)+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + ($routeParams.page-1)+'</a>');
        }

        if ($routeParams.page < res.data.PrevPage) {
          $scope.next = $routeParams.page + 1;
          // $('#next-page').html('<a href="#'+($routeParams.page+1)+'" id="next-page" class="page-link" onclick="location.reload()">' + ($routeParams.page+1)+' -></a>');
        } else {
          $scope.next = 0;
          // $('#next-page').html('<a href="#0" id="next-page" class="page-link" onclick="location.reload()">' + ($routeParams.page+1)+' last -></a>');
        }
      })
    }
  }]
});
