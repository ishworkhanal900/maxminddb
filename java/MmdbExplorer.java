import com.maxmind.geoip2.DatabaseReader;
import com.maxmind.geoip2.exception.GeoIp2Exception;
import com.maxmind.geoip2.model.CityResponse;
import com.maxmind.geoip2.record.City;

import java.io.File;
import java.io.IOException;
import java.net.InetAddress;
import java.util.Scanner;

public class MmdbExplorer {

    public static void main(String[] args) throws IOException, GeoIp2Exception {

        // Ask the user for the path to the MMDB file
        Scanner scanner = new Scanner(System.in);
        System.out.println("Enter the path to the MMDB file:");
        String filePath = scanner.nextLine();

        // Open the MMDB file
        File database = new File(filePath);
        DatabaseReader reader = new DatabaseReader.Builder(database).build();

        // Ask the user for the IP address to look up
        System.out.println("Enter the IP address to look up:");
        String ipAddressString = scanner.nextLine();

        // Look up the location data for the IP address
        InetAddress ipAddress = InetAddress.getByName(ipAddressString);
        CityResponse response = reader.city(ipAddress);
        City city = response.getCity();

        // Print the location information
        System.out.println("City: " + city.getName());
        System.out.println("Country: " + response.getCountry().getName());
        System.out.println("Latitude: " + response.getLocation().getLatitude());
        System.out.println("Longitude: " + response.getLocation().getLongitude());

        // Close the MMDB file
        reader.close();
    }
}

