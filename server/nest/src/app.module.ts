import { Module } from '@nestjs/common';
import { TypeOrmModule, TypeOrmModuleAsyncOptions } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TagModule } from './modules/tag/tag.module';
import { APP_FILTER, APP_INTERCEPTOR } from '@nestjs/core';
import { ResponseInterceptorInterceptor } from './interceptors/response-interceptor.interceptor';
import { HttpExceptionFilter } from './filters/http-exception.filter';
import { WinstonModule } from 'nest-winston';
import { AllExceptionFilter } from './filters/all-exception.filter';
import { ConfigModule, ConfigService } from '@nestjs/config';
import config from './config/config';
import * as winston from 'winston';
import 'winston-daily-rotate-file';

const format = winston.format;

@Module({
  imports: [
    TagModule,
    // config
    ConfigModule.forRoot({ isGlobal: true, load: [config], cache: true }),
    // database
    TypeOrmModule.forRootAsync({
      inject: [ConfigService],
      useFactory: async (configServer: ConfigService) => {
        const dbConfig = configServer.get('DataBase');
        return {
          type: dbConfig.DBType,
          host: dbConfig.Host,
          port: dbConfig.Port,
          username: dbConfig.Username,
          password: dbConfig.Password + '',
          database: dbConfig.DBName,
          entityPrefix: dbConfig.TablePrefix,
          entities: [__dirname + '/**/*/entities/*{.ts,.js}'],
          synchronize: false,
          logging: true,
        } as TypeOrmModuleAsyncOptions;
      },
    }),
    // logger
    WinstonModule.forRootAsync({
      inject: [ConfigService],
      useFactory: async (configServer: ConfigService) => {
        const logConfig = configServer.get('Log');

        const transport = new winston.transports.DailyRotateFile({
          level: 'error',
          filename: logConfig.LogFileName,
          dirname: logConfig.LogSavePath,
          datePattern: 'YY-MM-DD',
          maxSize: logConfig.MaxSize,
          maxFiles: logConfig.MaxFiles,
        });

        return {
          exitOnError: false,
          format: format.combine(
            format.timestamp({ format: 'YY-MM-DD hh:mm:ss A' }),
            format.json(),
            format.prettyPrint(),
          ),
          transports: [transport],
        };
      },
    }),
  ],
  controllers: [AppController],
  providers: [
    AppService,
    { provide: APP_INTERCEPTOR, useClass: ResponseInterceptorInterceptor },
    { provide: APP_FILTER, useClass: AllExceptionFilter },
    { provide: APP_FILTER, useClass: HttpExceptionFilter },
  ],
})
export class AppModule {}
