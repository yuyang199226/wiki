class Example {
    static void main(String[] args) {
        def name = "boss"
        float price = 20.99
        println(name)
        println(price)
        println("name=${name}, price=${price}")
        println('name=${name}, price=${price}')
        println 'hello World'
        println "Hello World";
        def ageMap = ["Ken": 22, "John": 25, "Sally": 22];
        for (am in ageMap) {
            println(am);
        }
    }
}