class BoutiqueInventory
  def initialize(items)
    @items = items
  end

  def item_names
    items.map { |x| x[:name] }.sort
  end

  def cheap
    items.select { |x| x[:price] < 30 }
  end

  def out_of_stock
    items.select { |x| stock_for_item(x[:name]).values.sum == 0 }
  end

  def stock_for_item(name)
    item = items.find { |x| x[:name] == name }
    
    item[:quantity_by_size] || 0
  end

  def total_stock
    items.map { |x| x[:quantity_by_size].values }.flatten.sum
  end

  private
  attr_reader :items
end
