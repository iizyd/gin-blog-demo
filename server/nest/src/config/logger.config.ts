import * as winston from 'winston';
import 'winston-daily-rotate-file';
const format = winston.format;

const transport = new winston.transports.DailyRotateFile({
  level: 'error',
  filename: 'application-%DATE%.log',
  dirname: 'src/storage/logs',
  datePattern: 'YY-MM-DD',
  maxSize: '7m',
  maxFiles: '7d',
});

export const loggerOption = {
  exitOnError: false,
  format: format.combine(
    format.timestamp({ format: 'YY-MM-DD hh:mm:ss A' }),
    format.json(),
    format.prettyPrint(),
  ),
  transports: [transport],
};
