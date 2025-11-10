/**
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {RouterModule, Routes} from '@angular/router';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {LoginComponent} from './login.component';
import {environment} from '../../../environments/environment';
import {initializeApp, provideFirebaseApp} from '@angular/fire/app';
import {
  getAnalytics,
  provideAnalytics,
  ScreenTrackingService,
  UserTrackingService,
} from '@angular/fire/analytics';
import {provideAuth, getAuth} from '@angular/fire/auth';

const routes: Routes = [{path: '', component: LoginComponent}];

const firebaseImports =
  environment.requiredLogin === 'True'
    ? [
        provideFirebaseApp(() => initializeApp(environment.firebase)),
        provideAuth(() => getAuth()),
        provideAnalytics(() => getAnalytics()),
      ]
    : [];

const firebaseProviders =
  environment.requiredLogin === 'True'
    ? [ScreenTrackingService, UserTrackingService]
    : [];

@NgModule({
  declarations: [LoginComponent],
  imports: [
    CommonModule,
    MatProgressSpinnerModule,
    RouterModule.forChild(routes),
    ...firebaseImports,
  ],
  providers: [...firebaseProviders],
})
export class LoginModule {}
