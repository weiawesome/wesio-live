export interface LogData {
  [key: string]: any;
}

export interface LogConfig {
  level: 'debug' | 'info' | 'warn' | 'error';
  format: 'json' | 'console';
}

interface LogEntry {
  level: string;
  trace_id?: string;
  component: string;
  action: string;
  message: string;
  data?: LogData;
  error?: string;
  timestamp: string;
}

export class Logger {
  private traceId?: string;
  private config: LogConfig;

  constructor(config: LogConfig, traceId?: string) {
    this.config = config;
    this.traceId = traceId;
  }

  /**
   * Creates a new logger with trace ID
   */
  withTraceId(traceId: string): Logger {
    return new Logger(this.config, traceId);
  }

  /**
   * Generates a new trace ID and returns logger with it
   */
  newTraceId(): Logger {
    const traceId = this.generateTraceId();
    return new Logger(this.config, traceId);
  }

  /**
   * Logs debug message with fixed JSON structure
   */
  debug(component: string, action: string, message: string, data?: LogData): void {
    if (this.shouldLog('debug')) {
      this.writeLog('debug', component, action, message, undefined, data);
    }
  }

  /**
   * Logs info message with fixed JSON structure
   */
  info(component: string, action: string, message: string, data?: LogData): void {
    if (this.shouldLog('info')) {
      this.writeLog('info', component, action, message, undefined, data);
    }
  }

  /**
   * Logs warning message with fixed JSON structure
   */
  warn(component: string, action: string, message: string, data?: LogData): void {
    if (this.shouldLog('warn')) {
      this.writeLog('warn', component, action, message, undefined, data);
    }
  }

  /**
   * Logs error message with fixed JSON structure
   */
  error(component: string, action: string, message: string, error?: Error, data?: LogData): void {
    if (this.shouldLog('error')) {
      this.writeLog('error', component, action, message, error, data);
    }
  }

  private shouldLog(level: string): boolean {
    const levels = ['debug', 'info', 'warn', 'error'];
    const configLevelIndex = levels.indexOf(this.config.level);
    const logLevelIndex = levels.indexOf(level);
    return logLevelIndex >= configLevelIndex;
  }

  private writeLog(
    level: string,
    component: string,
    action: string,
    message: string,
    error?: Error,
    data?: LogData
  ): void {
    const entry: LogEntry = {
      level,
      component,
      action,
      message,
      timestamp: new Date().toISOString(),
    };

    if (this.traceId) {
      entry.trace_id = this.traceId;
    }

    if (error) {
      entry.error = error.message;
    }

    if (data) {
      entry.data = data;
    }

    if (this.config.format === 'console') {
      this.writeConsoleLog(entry);
    } else {
      this.writeJsonLog(entry);
    }
  }

  private writeConsoleLog(entry: LogEntry): void {
    const timestamp = new Date(entry.timestamp).toLocaleTimeString();
    const traceInfo = entry.trace_id ? `[${entry.trace_id}] ` : '';
    const levelColor = this.getLevelColor(entry.level);
    
    console.log(
      `${timestamp} ${levelColor}${entry.level.toUpperCase()}${this.resetColor()} ${traceInfo}${entry.component}:${entry.action} - ${entry.message}`,
      entry.data ? entry.data : '',
      entry.error ? `Error: ${entry.error}` : ''
    );
  }

  private writeJsonLog(entry: LogEntry): void {
    console.log(JSON.stringify(entry));
  }

  private getLevelColor(level: string): string {
    if (typeof window !== 'undefined') return ''; // Browser environment
    
    const colors = {
      debug: '\x1b[36m', // Cyan
      info: '\x1b[32m',  // Green
      warn: '\x1b[33m',  // Yellow
      error: '\x1b[31m', // Red
    };
    return colors[level as keyof typeof colors] || '';
  }

  private resetColor(): string {
    if (typeof window !== 'undefined') return ''; // Browser environment
    return '\x1b[0m';
  }

  private generateTraceId(): string {
    const timestamp = new Date().toISOString().replace(/[-:.]/g, '').slice(0, 14);
    const random = this.randomString(8);
    return `${timestamp}-${random}`;
  }

  private randomString(length: number): string {
    const charset = 'abcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
      result += charset.charAt(Math.floor(Math.random() * charset.length));
    }
    return result;
  }
}

// Default configuration
export const defaultConfig: LogConfig = {
  level: 'info',
  format: 'json',
};

// Global logger instance
let globalLogger: Logger;

/**
 * Initializes the global logger with configuration
 */
export function init(config: LogConfig = defaultConfig): Logger {
  globalLogger = new Logger(config);
  return globalLogger;
}

/**
 * Gets the global logger instance
 */
export function getLogger(): Logger {
  if (!globalLogger) {
    globalLogger = new Logger(defaultConfig);
  }
  return globalLogger;
}

// Package level functions using global logger

/**
 * Creates a new logger with trace ID using global logger
 */
export function withTraceId(traceId: string): Logger {
  return getLogger().withTraceId(traceId);
}

/**
 * Generates a new trace ID using global logger
 */
export function newTraceId(): Logger {
  return getLogger().newTraceId();
}

/**
 * Logs debug message using global logger
 */
export function debug(component: string, action: string, message: string, data?: LogData): void {
  getLogger().debug(component, action, message, data);
}

/**
 * Logs info message using global logger
 */
export function info(component: string, action: string, message: string, data?: LogData): void {
  getLogger().info(component, action, message, data);
}

/**
 * Logs warning message using global logger
 */
export function warn(component: string, action: string, message: string, data?: LogData): void {
  getLogger().warn(component, action, message, data);
}

/**
 * Logs error message using global logger
 */
export function error(component: string, action: string, message: string, error?: Error, data?: LogData): void {
  getLogger().error(component, action, message, error, data);
}
