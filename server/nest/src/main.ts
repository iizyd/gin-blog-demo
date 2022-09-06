import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe, VersioningType } from '@nestjs/common';
import { WinstonModule } from 'nest-winston';
import * as winston from 'winston';
const format = winston.format;

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {
    logger: WinstonModule.createLogger({
      exitOnError: false,
      format: format.combine(
        format.colorize(),
        format.timestamp({
          format: 'HH:mm:ss YY/MM/DD',
        }),
        format.label({
          label: '测试',
        }),
        format.splat(),
        format.printf((info) => {
          return `${info.timestamp} ${info.level}: [${info.label}] ${info.message}`;
        }),
      ),
      transports: [
        new winston.transports.Console({
          level: 'info',
        }),
      ],
    }),
  });

  app.setGlobalPrefix('api');
  app.useGlobalPipes(new ValidationPipe({ transform: true, whitelist: true }));

  app.enableVersioning({
    defaultVersion: '1',
    type: VersioningType.URI,
  });

  await app.listen(8000);
}
bootstrap();
