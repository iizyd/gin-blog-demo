import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TagModule } from './modules/tag/tag.module';
import { APP_INTERCEPTOR } from '@nestjs/core';
import { ResponseInterceptorInterceptor } from './interceptors/response-interceptor.interceptor';

@Module({
  imports: [
    TagModule,
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'localhost',
      port: 3306,
      username: 'root',
      password: '123456',
      database: 'xigua_blog',
      entityPrefix: 'blog_',
      entities: [__dirname + '/**/*/entities/*{.ts,.js}'],
      synchronize: false,
    }),
  ],
  controllers: [AppController],
  providers: [
    AppService,
    { provide: APP_INTERCEPTOR, useClass: ResponseInterceptorInterceptor },
  ],
})
export class AppModule {}
