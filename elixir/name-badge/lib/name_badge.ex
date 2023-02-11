defmodule NameBadge do
  def print(id, name, department) do
    department = if department, do: department, else: "Owner"

    department_upcase = String.upcase(department)

    if id do
      "[#{id}] - #{name} - #{department_upcase}"
    else
      "#{name} - #{department_upcase}"
    end
  end
end
