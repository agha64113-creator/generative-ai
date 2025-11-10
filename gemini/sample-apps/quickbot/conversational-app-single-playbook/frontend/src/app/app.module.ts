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
import {BrowserModule} from '@angular/platform-browser';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HeaderComponent} from './components/header/header.component';
import {FooterComponent} from './components/footer/footer.component';
import {MainComponent} from './components/main/main.component';
import {ChatComponent} from './components/main/chat/chat.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatToolbarModule} from '@angular/material/toolbar';
import {HttpClientModule} from '@angular/common/http';
import {PathLocationStrategy} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {ChatbarComponent} from './components/main/chat/chatbar/chatbar.component';
import {MatTooltipModule} from '@angular/material/tooltip';
import {LoginComponent} from './components/login/login.component';
import {initializeApp, provideFirebaseApp} from '@angular/fire/app';
import {
  provideAnalytics,
  getAnalytics,
  ScreenTrackingService,
  UserTrackingService,
} from '@angular/fire/analytics';
import {environment} from '../environments/environment';
import {provideAuth, getAuth} from '@angular/fire/auth';
import {MatMenuModule} from '@angular/material/menu';
import {MatChipsModule} from '@angular/material/chips';
import {LocationStrategy} from '@angular/common';
import {MarkdownModule} from 'ngx-markdown';
import {MatDialogModule} from '@angular/material/dialog';
import {MatSnackBarModule} from '@angular/material/snack-bar';
import {SuggestionCardComponent} from './components/elements/suggestion-card/suggestion-card.component';
import {NgIdleModule} from '@ng-idle/core';
import {ClipboardModule} from '@angular/cdk/clipboard';

import {ToastMessageComponent} from './components/shared/toast-message/toast-message.component';
import {DialogueBoxComponent} from './dialogue-box/dialogue-box.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    MainComponent,
    ChatComponent,
    ChatbarComponent,
    LoginComponent,
    SuggestionCardComponent,
    ToastMessageComponent,
    DialogueBoxComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatSnackBarModule,
    MatToolbarModule,
    MatButtonModule,
    MatIconModule,
    HttpClientModule,
    FormsModule,
    MatTooltipModule,
    MatProgressSpinnerModule,
    MatMenuModule,
    MatChipsModule,
    ReactiveFormsModule,
    MatDialogModule,
    MarkdownModule.forRoot(),
    environment.requiredLogin === 'True'
      ? [
          provideFirebaseApp(() => initializeApp(environment.firebase)),
          provideAuth(() => getAuth()),
        ]
      : [],
    environment.requiredLogin === 'True'
      ? [provideAnalytics(() => getAnalytics())]
      : [],
    NgIdleModule.forRoot(),
    ClipboardModule,
  ],
  providers: [
    {provide: LocationStrategy, useClass: PathLocationStrategy},
    environment.requiredLogin === 'True'
      ? [
          ScreenTrackingService, // Automatically track screen views
          UserTrackingService, // Automatically track user interactions
        ]
      : [],
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
