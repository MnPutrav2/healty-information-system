import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '@/views/LoginView.vue'
import RegisterPatient from '@/components/Register/RegisterPatient.vue'
import FirstComponents from '@/components/FirstComponents.vue'
import NotFound from '@/components/NotFound.vue'
import Register from '@/components/Register/RegisterComponents.vue'
import Ambulatory from '@/components/AmbulatoryComponents.vue'
import Logs from '@/components/LogsComponents.vue'
import DrugDatas from '@/components/Pharmacy/DrugDatas.vue'
import RecipeInput from '@/components/Form/RecipeInput.vue'
import Recipes from '@/components/Pharmacy/RecipesComponents.vue'
import LabolatoriumData from '@/components/Labolatorium/LabolatoriumData.vue'
import LabolatoriumExamination from '@/components/Labolatorium/LabolatoriumExamination.vue'
import ExaminationCost from '@/components/Finance/ExaminationCost.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home-main',
      component: HomeView,
      children: [
        {
          path: '',
          name: 'home',
          component: FirstComponents
        },
        {
          path: 'register',
          name:'register',
          children: [
            {
              path: 'patient-registration',
              name: 'patient-registration',
              component: RegisterPatient
            },
            {
              path: 'service-registration',
              name: 'service-registration',
              component: Register
            },
          ]
        },
        {
          path: 'ambulatory-care',
          name: 'ambulatory-care',
          component: Ambulatory
        },
        {
          path: 'recipe-input',
          name: 'recipe-input',
          component: RecipeInput,
        },
        {
          path: 'logs',
          name: 'logs',
          component: Logs,
        },
        {
          path: 'pharmacy',
          name: 'pharmacy',
          children: [
            {
              path: '',
              name: 'drug-datas',
              component: DrugDatas
            },
            {
              path: 'drug-datas',
              name: 'drug-datas',
              component: DrugDatas
            },
            {
              path: 'recipes',
              name: 'recipes',
              component: Recipes
            },
          ]
        },
        {
          path: 'labolatorium',
          name: 'labolatorium',
          children: [
            {
              path: 'labolatorium-datas',
              name: 'labolatorium-datas',
              component: LabolatoriumData
            },
            {
              path: 'labolatorium-examination',
              name: 'labolatorium-examination',
              component: LabolatoriumExamination
            },
          ]
        },
        {
          path: 'finance',
          name:  'finance',
          children: [
            {
              path: 'examination-cost',
              name: 'examination-cost',
              component:  ExaminationCost
            }
          ]
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: NotFound
    }
  ]
})

export default router
