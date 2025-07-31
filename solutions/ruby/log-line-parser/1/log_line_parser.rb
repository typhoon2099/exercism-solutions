class LogLineParser
  def initialize(line)
    @line = line
  end

  LOG_REGEX = /\[(.*)\]: (.*)/

  def message
    matches[2].strip
  end

  def log_level
    matches[1].downcase
  end

  def reformat
    "#{message} (#{log_level})"
  end

  private

  def matches
    @matches ||= @line.match LOG_REGEX
  end
end
