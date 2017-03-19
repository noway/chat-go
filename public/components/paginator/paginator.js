'use strict';

angular.module('myApp.paginator', [])

.component('paginator', {
  templateUrl: 'components/paginator/paginator.html',
  controller: ['$routeParams', '$http', '$scope', function PaginatorController($routeParams, $http, $scope) {

    $scope.prev = NaN;
    $scope.next = NaN;
    
    this.$onInit = () => {
      console.log($routeParams);
      console.log($routeParams.page);
      var page = parseInt(this.page, 10)
      console.log(page);
      console.log(this.page);
      console.log(this);
      if (page == 0) {

        $http.get('/page').then(res => {
          if (res.data.PrevPage == page) {
            return;
          }
          // $('#prev-page').html('<a href="#'+res.data.PrevPage+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + res.data.PrevPage+'</a>');

          $scope.prev = res.data.PrevPage;
        })
      } else {
        $http.get('/page').then(res => {
          if (page > 1) {
            $scope.prev = page - 1;
            // $('#prev-page').html('<a href="#'+(page-1)+'" id="prev-page" class="page-link" onclick="location.reload()"><- ' + (page-1)+'</a>');
          }

          if (page < res.data.PrevPage) {
            $scope.next = page + 1;
            // $('#next-page').html('<a href="#'+(page+1)+'" id="next-page" class="page-link" onclick="location.reload()">' + (page+1)+' -></a>');
          } else {
            $scope.next = 0;
            // $('#next-page').html('<a href="#0" id="next-page" class="page-link" onclick="location.reload()">' + (page+1)+' last -></a>');
          }
        })
      }
    }
  }],
  bindings: {
    page: '<',
  }

});
