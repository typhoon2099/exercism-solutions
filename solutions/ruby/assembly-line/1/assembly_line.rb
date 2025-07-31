class AssemblyLine
  CARS_PER_HOUR = 221
  
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    CARS_PER_HOUR * successful_cars_per_hour * @speed
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end

  private

  def successful_cars_per_hour
    case @speed
      when 5..8
        return 0.9
      when 9
        return 0.8
      when 10
        return 0.77
    end

    return 1
  end
end
