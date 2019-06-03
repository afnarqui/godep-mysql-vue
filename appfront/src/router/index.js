import Vue from 'vue';
import Router from 'vue-router';
import Domain from '@/components/Domain';
import Domaincompare from '@/components/Domaincompare';
import Domainall from '@/components/Domainall';
Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/Domain',
      name: 'Domain',
      component: Domain
    },
    {
      path: '/Domaincompare',
      name: 'Domaincompare',
      component: Domaincompare
    },
    {
      path: '/',
      name: 'Domainall',
      component: Domainall
    }
  ]
});
